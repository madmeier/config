package config

import "log/slog"

const (
	defaultNameSpace = "default"
)

type NameSpaceConfig struct {
	// name of the namespace
	name string
}

func (n *NameSpaceConfig) Name() string {
	return n.name
}

func (a *NameSpaceConfig) NameSpace() string {
	return ""
}

func NewNameSpaceConfig(name string) *NameSpaceConfig {
	return &NameSpaceConfig{
		name: name,
	}
}

func (a *ArchConfig) CheckNameSpaces(log *slog.Logger) {
}

func getNameSpace(namespace ...string) string {
	ns := defaultNameSpace
	if len(namespace) > 0 {
		ns = namespace[0]
	}

	return ns
}
