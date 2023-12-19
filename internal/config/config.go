package config

import (
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/stackus/dotenv"

	"shopping/internal/rpc"
	"shopping/internal/web"
)

type (
	DatabaseConfig struct {
		Name string
		Conn string `required:"true" default:"postgresql://postgres:postgres@127.0.0.1:5432/shopping?sslmode=disable"`
	}

	NatsConfig struct {
		URL    string `required:"true" default:"127.0.0.1:4222"`
		Stream string `default:"shopping"`
	}

	AppConfig struct {
		Environment     string
		LogLevel        string `envconfig:"LOG_LEVEL" default:"DEBUG"`
		Database        DatabaseConfig
		Nats            NatsConfig
		Rpc             rpc.RpcConfig
		Web             web.WebConfig
		ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
	}
)

func InitConfig() (cfg AppConfig, err error) {
	if err = dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT"))); err != nil {
		return
	}

	err = envconfig.Process("", &cfg)

	return
}
