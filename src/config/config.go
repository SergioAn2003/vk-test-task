package config

import (
	"log"
	"os"

	"github.com/jinzhu/configor"
)

type Config struct {
	Test struct {
		TestString string `yaml:"test_string"`
	} `yaml:"test"`
}

func NewConfig(confPath string) (Config, error) {
	var c = Config{}
	err := configor.Load(&c, confPath)
	return c, err
}

func (c *Config) PostgresqlConnectionString() string {
	url := os.Getenv("PG_URL")
	if url == "" {
		log.Fatalln("отсутствует PG_URL")
	}

	return url
}

func (c *Config) RedisConnectionData() struct {
	Host     string
	Password string
} {
	host := os.Getenv("REDIS_HOST")
	password := os.Getenv("REDIS_PASSWORD")

	if host == "" || password == "" {
		log.Fatalln("отсутствуют данные для подключения к redis")
	}
	return struct {
		Host     string
		Password string
	}{
		Host:     host,
		Password: password,
	}
}
