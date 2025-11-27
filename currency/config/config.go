package config

type MysqlConfig struct {
	Port         string `koanf:"port"`
	DatabaseName string `koanf:"database_name"`
	DatabasePass string `koanf:"database_pass"`
	Host         string `koanf:"host"`
}

type UserConfig struct {
	MysqlConfig MysqlConfig `koanf:"mysql_config"`
}
