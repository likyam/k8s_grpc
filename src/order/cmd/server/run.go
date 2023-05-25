package server

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"likyam.cn/src/common/config"
	"likyam.cn/src/common/server"
	serverV1 "likyam.cn/src/order/service/v1/dao"
	"net"
	"net/http"
	"os"
	"syscall"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/oklog/run"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"

	orderPBV1 "likyam.cn/api/gen/go/proto/order/v1"
	"likyam.cn/src/order/service/model"
)

// Server 启动微服务所需要的所有依赖
type Server struct {
	repo       *serverV1.Repository
	conf       *config.Config
	g          *run.Group
	logger     log.Logger
	httpServer *http.Server
	grpcServer *grpc.Server
	db         *gorm.DB
	trace      *sdktrace.TracerProvider
}

// NewServer 实例化 Server
func NewServer(
	repo *serverV1.Repository,
	conf *config.Config,
	g *run.Group,
	logger log.Logger,
	httpServer *http.Server,
	grpcServer *grpc.Server,
	db *gorm.DB,
	trace *sdktrace.TracerProvider,
) *Server {
	return &Server{
		repo:       repo,
		conf:       conf,
		g:          g,
		logger:     logger,
		httpServer: httpServer,
		grpcServer: grpcServer,
		db:         db,
		trace:      trace,
	}
}

// NewGrpcServer 实例化 Grpc 服务
func NewGrpcServer(orderServer orderPBV1.OrderServiceServer) *grpc.Server {

	grpcServer := grpc.NewServer(
		grpc.ChainStreamInterceptor(
			// otel 链路追踪
			otelgrpc.StreamServerInterceptor(),
			// 鉴权中间件
			auth.StreamServerInterceptor(server.AuthenticationInterceptor),
		),
		grpc.ChainUnaryInterceptor(
			// otel 链路追踪
			otelgrpc.UnaryServerInterceptor(),
			// 鉴权中间件
			auth.UnaryServerInterceptor(server.AuthenticationInterceptor),
			// PGV 中间件
			server.ValidationUnaryInterceptor,
		),
	)
	orderPBV1.RegisterOrderServiceServer(grpcServer, orderServer)
	return grpcServer

}

// NewHttpServer 实例化 Http 服务
func NewHttpServer(conf *config.Config) *http.Server {

	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(config.HttpErrorHandler),
		runtime.WithForwardResponseOption(config.HttpSuccessResponseModifier),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &config.CustomMarshaller{}),
	)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	if err := orderPBV1.RegisterOrderServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		conf.Grpc.Port,
		opts,
	); err != nil {
		panic("register service handler failed.[ERROR]=>" + err.Error())
	}

	httpServer := &http.Server{
		Addr:    conf.Http.Port,
		Handler: mux,
	}
	return httpServer

}

// RunServer 启动 http 以及 grpc 服务
func (s *Server) RunServer() {

	// 执行 migrate
	model.Migrate(s.db)

	// 启动 grpc 服务
	s.g.Add(func() error {
		l, err := net.Listen("tcp", s.conf.Grpc.Port)
		if err != nil {
			return err
		}
		_ = level.Info(s.logger).Log("msg", "starting gRPC server", "addr", l.Addr().String())
		return s.grpcServer.Serve(l)
	}, func(err error) {
		s.grpcServer.GracefulStop()
		s.grpcServer.Stop()
	})

	// 启动 http 服务
	s.g.Add(func() error {

		_ = level.Info(s.logger).Log("msg", "starting HTTP server", "addr", s.httpServer.Addr)
		return s.httpServer.ListenAndServe()

	}, func(err error) {
		if err = s.httpServer.Close(); err != nil {
			_ = level.Error(s.logger).Log("msg", "failed to stop web server", "err", err)
		}
	})
	s.g.Add(run.SignalHandler(context.Background(), syscall.SIGINT, syscall.SIGTERM))

	if err := s.g.Run(); err != nil {
		_ = level.Error(s.logger).Log("err", err)
		os.Exit(1)
	}

}

// Run 启动服务
func Run(cfg string) {

	initServer, err := InitServer(cfg)
	if err != nil {
		panic("run initServer failed.[ERROR]=>" + err.Error())
	}

	// 上报链路 trace 数据
	defer func() {
		if err = initServer.trace.Shutdown(context.Background()); err != nil {
			_ = level.Info(initServer.logger).Log("msg", "shutdown trace provider failed", "err", err)
		}
	}()

	initServer.RunServer()

}
