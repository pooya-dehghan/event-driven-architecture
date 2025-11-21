package config

type HttpServer struct {
	Port string `koanf: "port"`
}

type MysqlDatabase struct {
	Port         string `koanf: "port"`
	DatabaseName string `koanf: "database_name"`
	DatabasePass string `koanf: "database_pass"`
	Host         string `koanf: "host"`
}

type UserConfig struct {
	UserMicroservicePort HttpServer    `koanf:"http_server"`
	MysqlDatabase        MysqlDatabase `koanf:"mysql_database"`
}
