package sample

import (
	"log/slog"

	"github.com/blueorb/config/config"
)

const (
	ArchitectureName = "blueorb"
)

func DefineArchitecture(log *slog.Logger) (arch *config.ArchConfig) {
	// create the architecture
	arch = config.NewArchConfig(ArchitectureName)

	// define namespaces
	arch.RegisterNameSpace(log, NamespaceBlueOrb(log))

	// define secrets
	arch.RegisterSecret(log, SecretsDatabase(log))

	// define re-usable config maps
	arch.RegisterConfigMap(log, GoogleProjectMap(log))
	arch.RegisterConfigMap(log, DatabaseMap(log))
	arch.RegisterConfigMap(log, TLSMap(log))

	// define environments
	arch.RegisterEnvironment(log, SandboxEnvironment(log))
	arch.RegisterEnvironment(log, ProductionEnvironment(log))

	// define services
	arch.RegisterService(log, MicroService(log))

	return
}
