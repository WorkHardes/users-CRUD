package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

const (
	defaultHTTPPort           = "8000"
	defaultHTTPReadTimeout    = 15 * time.Second
	defaultHTTPWriteTimeout   = 15 * time.Second
	defaultHTTPMaxHeaderBytes = 1024
)

type Config struct {
	HTTP     HTTPConfig
	Postgres PostgresConfig
}

type HTTPConfig struct {
	Host           string
	Port           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

type PostgresConfig struct {
	UserName       string
	Password       string
	Host           string
	Port           string
	DBName         string
	MigrationsPath string
}

func Init(configsDir string) (*Config, error) {
	populateDefaults()

	if err := parseConfigsFile(configsDir); err != nil {
		return nil, fmt.Errorf("internal.config.parseConfigsFile failed\n%w;", err)
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("internal.config.unmarshal failed\n%w;", err)
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return fmt.Errorf("viper.UnmarshalKey failed\n%w;", err)
	}

	return nil
}

func setFromEnv(cfg *Config) {
	cfg.HTTP.Host = os.Getenv("HTTP_HOST")
	cfg.Postgres.UserName = os.Getenv("DB_USERNAME")
	cfg.Postgres.Password = os.Getenv("DB_PASSWORD")
	cfg.Postgres.Host = os.Getenv("DB_HOST")
	cfg.Postgres.Port = os.Getenv("DB_PORT")
	cfg.Postgres.DBName = os.Getenv("DB_NAME")
	cfg.Postgres.MigrationsPath = os.Getenv("DB_MIGRATIONS_PATH")
}

func parseConfigsFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("viper.ReadInConfig failed\n%w;", err)
	}

	return viper.MergeInConfig()
}

func populateDefaults() {
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.max_header_megabytes", defaultHTTPMaxHeaderBytes)
	viper.SetDefault("http.timeouts.read", defaultHTTPReadTimeout)
	viper.SetDefault("http.timeouts.write", defaultHTTPWriteTimeout)
}
