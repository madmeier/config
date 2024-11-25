package config

type NamedConfig interface {
	Name() string
	NameSpace() string
}
