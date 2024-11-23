package sample

import (
	"log/slog"

	"github.com/blueorb/config/config"
)

func defineArchitecture(log *slog.Logger) (arch *config.ArchConfig) {
	// create the architecture
	arch = config.NewArchConfig("blueorb")

	// define namespaces
	arch.RegisterNameSpace(log, NamespaceBlueOrb(log))

	// define environments
	arch.RegisterEnvironment(log, SandboxEnvironment(log))
	arch.RegisterEnvironment(log, ProductionEnvironment(log))

	// define services
	arch.RegisterService(log, MicroService(log))

	return
}
