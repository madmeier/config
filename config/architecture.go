package config

import (
	"fmt"
	"log/slog"
)

type (
	mapType[E NamedConfig] map[string]E
	nameSpaceMap           mapType[*NameSpaceConfig]
	serviceMap             mapType[*SvcConfig]
	environmentMap         mapType[*EnvConfig]
)

type ArchConfig struct {
	// name of the architecture
	name string
	// name spaces in the architecture
	nameSpaces map[string]*NameSpaceConfig
	// environments of the architecture
	environments map[string]*EnvConfig
	// services in the architecture
	services map[string]*SvcConfig

	// collect errors during architecture definition
	errors []error
}

func (a *ArchConfig) Name() string {
	return a.name
}

func (a *ArchConfig) Environments() environmentMap {
	return a.environments
}

func (a *ArchConfig) NameSpaces() []NameSpaceConfig {
	return a.NameSpaces()
}

func (a *ArchConfig) Services() serviceMap {
	return a.services
}

func (a *ArchConfig) Errors() []error {
	return a.errors
}

func registerNamedConfig[T NamedConfig](log *slog.Logger, a *ArchConfig, configs mapType[T], config T) (err error) {
	name := config.Name()
	if _, ok := configs[name]; ok {
		err = fmt.Errorf("%s already exists", name)
		a.AddError(err)
		log.Error("Registration failed", "error", err)
		return
	}

	configs[name] = config
	return
}

func (a *ArchConfig) RegisterNameSpace(log *slog.Logger, ns *NameSpaceConfig) (err error) {
	return registerNamedConfig(log, a, a.nameSpaces, ns)
}

func (a *ArchConfig) RegisterEnvironment(log *slog.Logger, env *EnvConfig) (err error) {
	return registerNamedConfig(log, a, a.environments, env)
}

func (a *ArchConfig) RegisterService(log *slog.Logger, svc *SvcConfig) (err error) {
	return registerNamedConfig(log, a, a.services, svc)
}

func NewArchConfig(name string) *ArchConfig {
	return &ArchConfig{
		name: name,
	}
}

func (a *ArchConfig) AddError(err error) {
	a.errors = append(a.errors, err)
}
