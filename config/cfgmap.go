package config

import "log/slog"

type ValueOptions struct {
	SecretName   string
	SecretKey    string
	DefaultValue any
}

type ValueOption func(*ValueOptions)

type ConfigMap struct {
	name      string
	nameSpace string
	// environment variables with as generic default value
	properties map[string]ValueOptions
}

func NewConfigMap(name string, namespace ...string) *ConfigMap {
	return &ConfigMap{
		name:       name,
		nameSpace:  getNameSpace(namespace...),
		properties: map[string]ValueOptions{},
	}
}

func (m *ConfigMap) Name() string {
	return m.name
}

func (m *ConfigMap) NameSpace() string {
	return m.nameSpace
}

func (m *ConfigMap) Properties() map[string]ValueOptions {
	return m.properties
}

func extractValueOptions(
	options ...ValueOption,
) (
	valueOptions ValueOptions,
) {
	valueOptions = ValueOptions{}
	for _, opt := range options {
		opt(&valueOptions)
	}

	return
}

func (m *ConfigMap) Add(key string, options ...ValueOption) {
	m.properties[key] = extractValueOptions(options...)
}

func WithDefaultValue(defaultValue any) ValueOption {
	return func(s *ValueOptions) {
		s.DefaultValue = defaultValue
	}
}

func WithSecretRef(secretName string, secretKey string) ValueOption {
	return func(s *ValueOptions) {
		s.SecretName = secretName
		s.SecretKey = secretKey
	}
}

func (a *ArchConfig) CheckConfigMaps(log *slog.Logger) {
}
