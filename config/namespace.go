package config

type NameSpaceConfig struct {
	// name of the namespace
	name string
}

func (n *NameSpaceConfig) Name() string {
	return n.name
}

func NewNameSpaceConfig(name string) *NameSpaceConfig {
	return &NameSpaceConfig{
		name: name,
	}
}
