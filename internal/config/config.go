package config

import (
	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigyaml"
	"log"
	"sync"
)

type Config struct {
	Port        string `yaml:"port" env:"PORT" required:"true"`
	DatabaseDSN string `yaml:"database_dsn" env:"DATABASE_DSN" default:"postgres://postgres:postgres@localhost:5432/book_db?sslmode=disable"`
}

var (
	cfg  Config
	once sync.Once
)

func Get() Config {
	once.Do(func() {
		loader := aconfig.LoaderFor(&cfg, aconfig.Config{
			EnvPrefix: "NFB",
			Files:     []string{"./config.yaml", "./config.local.yaml"},
			FileDecoders: map[string]aconfig.FileDecoder{
				".yaml": aconfigyaml.New(),
			},
		})

		if err := loader.Load(); err != nil {
			log.Printf("[ERROR] failed to load config: %v", err)
		}
	})

	return cfg
}
