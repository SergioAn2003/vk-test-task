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

func (c *Config) PostgresConnectionString() string {
	url := os.Getenv("PG_URL")
	if url == "" {
		log.Fatalln("отсутствует PG_URL")
	}

	return url
}
