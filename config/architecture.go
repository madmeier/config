package config

import (
	"fmt"
	"log/slog"
)

type (
	mapType[E ArchElement] map[string]E
	nameSpaceMap           mapType[*NameSpaceConfig]
	serviceMap             mapType[*SvcConfig]
	environmentMap         mapType[*EnvConfig]
	secretsMap             mapType[*Secret]
	configMap              mapType[*ConfigMap]
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
	// config maps in the architecture
	configMaps map[string]*ConfigMap
	// secrets in the architecture
	secrets map[string]*Secret

	// collect errors during architecture definition
	errors []error
}

func (a *ArchConfig) Name() string {
	return a.name
}

func (a *ArchConfig) Environments() environmentMap {
	return a.environments
}

func (a *ArchConfig) NameSpaces() nameSpaceMap {
	return a.nameSpaces
}

func (a *ArchConfig) ConfigMaps() configMap {
	return a.configMaps
}

func (a *ArchConfig) Secrets() secretsMap {
	return a.secrets
}

func (a *ArchConfig) Services() serviceMap {
	return a.services
}

func (a *ArchConfig) Errors() []error {
	return a.errors
}

func registerNamedConfig[T ArchElement](log *slog.Logger, a *ArchConfig, configs mapType[T], config T) (err error) {
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

func (a *ArchConfig) RegisterConfigMap(log *slog.Logger, m *ConfigMap) (err error) {
	return registerNamedConfig(log, a, a.configMaps, m)
}

func (a *ArchConfig) RegisterSecret(log *slog.Logger, s *Secret) (err error) {
	return registerNamedConfig(log, a, a.secrets, s)
}

func NewArchConfig(name string) *ArchConfig {
	return &ArchConfig{
		name:         name,
		nameSpaces:   make(map[string]*NameSpaceConfig),
		environments: make(map[string]*EnvConfig),
		services:     make(map[string]*SvcConfig),
		configMaps:   make(map[string]*ConfigMap),
		secrets:      make(map[string]*Secret),
	}
}

func (a *ArchConfig) AddError(err error) {
	a.errors = append(a.errors, err)
}

func (a *ArchConfig) ResetErrors() {
	a.errors = []error{}
}

// check the consistency of the architecture
func (a *ArchConfig) Check(log *slog.Logger) (err error) {
	a.ResetErrors()

	a.CheckNameSpaces(log)
	a.CheckEnvironments(log)
	a.CheckConfigMaps(log)
	a.CheckSecrets(log)
	a.CheckServices(log)

	if len(a.errors) > 0 {
		err = fmt.Errorf("architecture definition contains errors")
		log.Error("Architecture definition contains errors", "errors", a.errors)
	}

	return
}
