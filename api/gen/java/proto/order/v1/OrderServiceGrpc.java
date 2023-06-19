package order.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.56.0)",
    comments = "Source: proto/order/v1/order.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class OrderServiceGrpc {

  private OrderServiceGrpc() {}

  public static final String SERVICE_NAME = "order.v1.OrderService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<order.v1.OrderOuterClass.CreateOrderRequest,
      order.v1.OrderOuterClass.CreateOrderResponse> getCreateOrderMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateOrder",
      requestType = order.v1.OrderOuterClass.CreateOrderRequest.class,
      responseType = order.v1.OrderOuterClass.CreateOrderResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<order.v1.OrderOuterClass.CreateOrderRequest,
      order.v1.OrderOuterClass.CreateOrderResponse> getCreateOrderMethod() {
    io.grpc.MethodDescriptor<order.v1.OrderOuterClass.CreateOrderRequest, order.v1.OrderOuterClass.CreateOrderResponse> getCreateOrderMethod;
    if ((getCreateOrderMethod = OrderServiceGrpc.getCreateOrderMethod) == null) {
      synchronized (OrderServiceGrpc.class) {
        if ((getCreateOrderMethod = OrderServiceGrpc.getCreateOrderMethod) == null) {
          OrderServiceGrpc.getCreateOrderMethod = getCreateOrderMethod =
              io.grpc.MethodDescriptor.<order.v1.OrderOuterClass.CreateOrderRequest, order.v1.OrderOuterClass.CreateOrderResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateOrder"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  order.v1.OrderOuterClass.CreateOrderRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  order.v1.OrderOuterClass.CreateOrderResponse.getDefaultInstance()))
              .setSchemaDescriptor(new OrderServiceMethodDescriptorSupplier("CreateOrder"))
              .build();
        }
      }
    }
    return getCreateOrderMethod;
  }

  private static volatile io.grpc.MethodDescriptor<order.v1.OrderOuterClass.GetOrderByIdRequest,
      order.v1.OrderOuterClass.GetOrderByIdResponse> getGetOrderByIdMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetOrderById",
      requestType = order.v1.OrderOuterClass.GetOrderByIdRequest.class,
      responseType = order.v1.OrderOuterClass.GetOrderByIdResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<order.v1.OrderOuterClass.GetOrderByIdRequest,
      order.v1.OrderOuterClass.GetOrderByIdResponse> getGetOrderByIdMethod() {
    io.grpc.MethodDescriptor<order.v1.OrderOuterClass.GetOrderByIdRequest, order.v1.OrderOuterClass.GetOrderByIdResponse> getGetOrderByIdMethod;
    if ((getGetOrderByIdMethod = OrderServiceGrpc.getGetOrderByIdMethod) == null) {
      synchronized (OrderServiceGrpc.class) {
        if ((getGetOrderByIdMethod = OrderServiceGrpc.getGetOrderByIdMethod) == null) {
          OrderServiceGrpc.getGetOrderByIdMethod = getGetOrderByIdMethod =
              io.grpc.MethodDescriptor.<order.v1.OrderOuterClass.GetOrderByIdRequest, order.v1.OrderOuterClass.GetOrderByIdResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetOrderById"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  order.v1.OrderOuterClass.GetOrderByIdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  order.v1.OrderOuterClass.GetOrderByIdResponse.getDefaultInstance()))
              .setSchemaDescriptor(new OrderServiceMethodDescriptorSupplier("GetOrderById"))
              .build();
        }
      }
    }
    return getGetOrderByIdMethod;
  }

  private static volatile io.grpc.MethodDescriptor<order.v1.OrderOuterClass.healthyRequest,
      order.v1.OrderOuterClass.healthyResponse> getHealthyMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Healthy",
      requestType = order.v1.OrderOuterClass.healthyRequest.class,
      responseType = order.v1.OrderOuterClass.healthyResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<order.v1.OrderOuterClass.healthyRequest,
      order.v1.OrderOuterClass.healthyResponse> getHealthyMethod() {
    io.grpc.MethodDescriptor<order.v1.OrderOuterClass.healthyRequest, order.v1.OrderOuterClass.healthyResponse> getHealthyMethod;
    if ((getHealthyMethod = OrderServiceGrpc.getHealthyMethod) == null) {
      synchronized (OrderServiceGrpc.class) {
        if ((getHealthyMethod = OrderServiceGrpc.getHealthyMethod) == null) {
          OrderServiceGrpc.getHealthyMethod = getHealthyMethod =
              io.grpc.MethodDescriptor.<order.v1.OrderOuterClass.healthyRequest, order.v1.OrderOuterClass.healthyResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Healthy"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  order.v1.OrderOuterClass.healthyRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  order.v1.OrderOuterClass.healthyResponse.getDefaultInstance()))
              .setSchemaDescriptor(new OrderServiceMethodDescriptorSupplier("Healthy"))
              .build();
        }
      }
    }
    return getHealthyMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static OrderServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<OrderServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<OrderServiceStub>() {
        @java.lang.Override
        public OrderServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new OrderServiceStub(channel, callOptions);
        }
      };
    return OrderServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static OrderServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<OrderServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<OrderServiceBlockingStub>() {
        @java.lang.Override
        public OrderServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new OrderServiceBlockingStub(channel, callOptions);
        }
      };
    return OrderServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static OrderServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<OrderServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<OrderServiceFutureStub>() {
        @java.lang.Override
        public OrderServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new OrderServiceFutureStub(channel, callOptions);
        }
      };
    return OrderServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void createOrder(order.v1.OrderOuterClass.CreateOrderRequest request,
        io.grpc.stub.StreamObserver<order.v1.OrderOuterClass.CreateOrderResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateOrderMethod(), responseObserver);
    }

    /**
     */
    default void getOrderById(order.v1.OrderOuterClass.GetOrderByIdRequest request,
        io.grpc.stub.StreamObserver<order.v1.OrderOuterClass.GetOrderByIdResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetOrderByIdMethod(), responseObserver);
    }

    /**
     * <pre>
     * Healthy
     * </pre>
     */
    default void healthy(order.v1.OrderOuterClass.healthyRequest request,
        io.grpc.stub.StreamObserver<order.v1.OrderOuterClass.healthyResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getHealthyMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service OrderService.
   */
  public static abstract class OrderServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return OrderServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service OrderService.
   */
  public static final class OrderServiceStub
      extends io.grpc.stub.AbstractAsyncStub<OrderServiceStub> {
    private OrderServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected OrderServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new OrderServiceStub(channel, callOptions);
    }

    /**
     */
    public void createOrder(order.v1.OrderOuterClass.CreateOrderRequest request,
        io.grpc.stub.StreamObserver<order.v1.OrderOuterClass.CreateOrderResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateOrderMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getOrderById(order.v1.OrderOuterClass.GetOrderByIdRequest request,
        io.grpc.stub.StreamObserver<order.v1.OrderOuterClass.GetOrderByIdResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetOrderByIdMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Healthy
     * </pre>
     */
    public void healthy(order.v1.OrderOuterClass.healthyRequest request,
        io.grpc.stub.StreamObserver<order.v1.OrderOuterClass.healthyResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getHealthyMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service OrderService.
   */
  public static final class OrderServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<OrderServiceBlockingStub> {
    private OrderServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected OrderServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new OrderServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public order.v1.OrderOuterClass.CreateOrderResponse createOrder(order.v1.OrderOuterClass.CreateOrderRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateOrderMethod(), getCallOptions(), request);
    }

    /**
     */
    public order.v1.OrderOuterClass.GetOrderByIdResponse getOrderById(order.v1.OrderOuterClass.GetOrderByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetOrderByIdMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Healthy
     * </pre>
     */
    public order.v1.OrderOuterClass.healthyResponse healthy(order.v1.OrderOuterClass.healthyRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getHealthyMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service OrderService.
   */
  public static final class OrderServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<OrderServiceFutureStub> {
    private OrderServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected OrderServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new OrderServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<order.v1.OrderOuterClass.CreateOrderResponse> createOrder(
        order.v1.OrderOuterClass.CreateOrderRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateOrderMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<order.v1.OrderOuterClass.GetOrderByIdResponse> getOrderById(
        order.v1.OrderOuterClass.GetOrderByIdRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetOrderByIdMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Healthy
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<order.v1.OrderOuterClass.healthyResponse> healthy(
        order.v1.OrderOuterClass.healthyRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getHealthyMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_ORDER = 0;
  private static final int METHODID_GET_ORDER_BY_ID = 1;
  private static final int METHODID_HEALTHY = 2;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_CREATE_ORDER:
          serviceImpl.createOrder((order.v1.OrderOuterClass.CreateOrderRequest) request,
              (io.grpc.stub.StreamObserver<order.v1.OrderOuterClass.CreateOrderResponse>) responseObserver);
          break;
        case METHODID_GET_ORDER_BY_ID:
          serviceImpl.getOrderById((order.v1.OrderOuterClass.GetOrderByIdRequest) request,
              (io.grpc.stub.StreamObserver<order.v1.OrderOuterClass.GetOrderByIdResponse>) responseObserver);
          break;
        case METHODID_HEALTHY:
          serviceImpl.healthy((order.v1.OrderOuterClass.healthyRequest) request,
              (io.grpc.stub.StreamObserver<order.v1.OrderOuterClass.healthyResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getCreateOrderMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              order.v1.OrderOuterClass.CreateOrderRequest,
              order.v1.OrderOuterClass.CreateOrderResponse>(
                service, METHODID_CREATE_ORDER)))
        .addMethod(
          getGetOrderByIdMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              order.v1.OrderOuterClass.GetOrderByIdRequest,
              order.v1.OrderOuterClass.GetOrderByIdResponse>(
                service, METHODID_GET_ORDER_BY_ID)))
        .addMethod(
          getHealthyMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              order.v1.OrderOuterClass.healthyRequest,
              order.v1.OrderOuterClass.healthyResponse>(
                service, METHODID_HEALTHY)))
        .build();
  }

  private static abstract class OrderServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    OrderServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return order.v1.OrderOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("OrderService");
    }
  }

  private static final class OrderServiceFileDescriptorSupplier
      extends OrderServiceBaseDescriptorSupplier {
    OrderServiceFileDescriptorSupplier() {}
  }

  private static final class OrderServiceMethodDescriptorSupplier
      extends OrderServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    OrderServiceMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (OrderServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new OrderServiceFileDescriptorSupplier())
              .addMethod(getCreateOrderMethod())
              .addMethod(getGetOrderByIdMethod())
              .addMethod(getHealthyMethod())
              .build();
        }
      }
    }
    return result;
  }
}
