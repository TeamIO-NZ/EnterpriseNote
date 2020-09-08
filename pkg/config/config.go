package config

type Config struct {
	webserver  WebServerConfig  `json:"webserver"`
	postgresql PostgresqlConfig `json:"postgresql"`
}

type WebServerConfig struct {
	port int `json:"port"`
}

type PostgresqlConfig struct {
	address string `json:"address"`
	port    int    `json:"port"`
}
