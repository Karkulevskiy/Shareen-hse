package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config is a description for application
type Config struct {
	Env              string `yaml:"env" env:"ENV" env-default:"local"`
	ConnectionString string `yaml:"connection_string"`
	HTTPServer       `yaml:"http_server"`
}

// HTTPServer is a description for http server
type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

// Initializing configuration
func MustLoad() *Config {
	cfgPath := fetchConfigPath()

	if cfgPath == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		panic("config file not found:" + cfgPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

// Fetching config file from --flag or EnvVariable
func fetchConfigPath() string {
	var cfgPath string

	// --config="path/to/config.yaml"
	flag.StringVar(&cfgPath, "config", "", "path to configuration file")

	flag.Parse()

	if cfgPath == "" {
		cfgPath = os.Getenv("CONFIG_PATH")
	}

	return cfgPath
}
