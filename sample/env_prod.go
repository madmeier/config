package sample

import (
	"log/slog"

	"github.com/blueorb/config/config"
)

const (
	EnvProdName = "prod"
)

func ProductionEnvironment(log *slog.Logger) (env *config.EnvConfig) {
	env = config.NewEnvConfig(
		"EnvProdName",
		"The prod environment is used by clients",
	)
	return
}
