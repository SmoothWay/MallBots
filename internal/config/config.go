package config

import (
	"os"
	"time"

	"github.com/SmoothWay/MallBots/internal/rpc"
	"github.com/SmoothWay/MallBots/internal/web"
	"github.com/kelseyhightower/envconfig"

	"github.com/stackus/dotenv"
)

type (
	PGConfig struct {
		Conn string `required:"true"`
	}

	AppConfig struct {
		Environment     string
		LogLevel        string `envconfig:"LOG_LEVEL" default:"DEBUG"`
		PG              PGConfig
		Rpc             rpc.RpcConfig
		Web             web.WebConfig
		ShutDownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
	}
)

func InitConfig() (*AppConfig, error) {
	if err := dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT"))); err != nil {
		return nil, err
	}
	var cfg AppConfig
	err := envconfig.Process("", &cfg)

	return &cfg, err
}
