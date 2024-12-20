package config

// ServerConfig 服务启动端口号配置
type ServerConfig struct {
	Port int `mapstructure:"port"`
}
