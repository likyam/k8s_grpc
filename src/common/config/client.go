package config

// Client Grpc 客户端配置
type Client struct {
	// auth 服务客户端配置
	UserHost string `json:"userHost" yaml:"userHost"`
	UserPort string `json:"userPort" yaml:"userPort"`

	OrderHost string `json:"orderHost" yaml:"orderHost"`
	OrderPort string `json:"orderPort" yaml:"orderPort"`
}
