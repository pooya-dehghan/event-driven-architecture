package config

type HttpServer struct {
	Port string `koanf: "port"`
}

type UserConfig struct {
	UserMicroservicePort HttpServer `koanf:"http_server"`
}
