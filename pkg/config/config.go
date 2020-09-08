package config

type Config struct {
	WebServer  WebServerConfig  `json:"webserver"`
	Postgresql PostgresqlConfig `json:"postgresql"`
}

type WebServerConfig struct {
	Port string `json:"port"`
}

type PostgresqlConfig struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

func (c Config) Init(filename string) (err error) {

	return nil
}

func (c Config) Validate() (err error) {
	return nil
}
