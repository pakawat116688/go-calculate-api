package config

type Config struct {
	Server struct {
		Port string
	}
	Redis RedisConfig
}

type RedisConfig struct {
	Host string
	Port string
}
