package config

type EnvConfig struct {
	// name of the environment
	name string
	// the purpose
	purpose string
}

func NewEnvConfig(name string, purpose string) *EnvConfig {
	return &EnvConfig{
		name:    name,
		purpose: purpose,
	}
}

func (e *EnvConfig) Name() string {
	return e.name
}

func (e *EnvConfig) Purpose() string {
	return e.purpose
}
