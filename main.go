package main

import (
	"fmt"
	"strings"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/pakawatkung/go-calculate-api/config"
	"github.com/pakawatkung/go-calculate-api/handler"
	"github.com/spf13/viper"
)

func main() {

	cfg, err := initConfig()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v %v %v", cfg.Server.Port, cfg.Redis.Host, cfg.Redis.Port)

	redisClint := initRedis(*cfg)

	app := fiber.New()

	app.Get("/calculate", handler.NewHandler(handler.NewCheckOperator(), handler.NewGetRedis(redisClint), handler.NewSetRedis(redisClint)))

	app.Listen(cfg.Server.Port)

}

func initConfig() (*config.Config, error) {

	cfg := config.Config{}

	viper.SetDefault("Server.Port", ":8000")
	viper.SetDefault("Redis.Host", "localhost")
	viper.SetDefault("Redis.Port", "6379")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func initRedis(cfg config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", cfg.Redis.Host, cfg.Redis.Port),
	})
}
