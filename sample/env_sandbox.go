package sample

import (
	"log/slog"

	"github.com/blueorb/config/config"
)

const (
	EnvSandboxName = "sandbox"
)

func SandboxEnvironment(log *slog.Logger) (env *config.EnvConfig) {
	env = config.NewEnvConfig(
		"EnvSandboxName",
		"The sandbox is ideal for developers to test and integrate new features",
	)
	return
}
