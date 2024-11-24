package config

type ValueOptions struct {
	SecretName   string
	SecretKey    string
	DefaultValue string
}

type ValueOption func(*ValueOptions)

type Map struct {
	name string
	// environment variables with as generic default value
	properties map[string]ValueOptions
}

func NewMap(name string) *Map {
	return &Map{
		name:       name,
		properties: map[string]ValueOptions{},
	}
}

func (m *Map) Name() string {
	return m.name
}

func (m *Map) Properties() map[string]ValueOptions {
	return m.properties
}

func extractValueOptions(
	key string,
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

func (m *Map) Add(key string, options ...ValueOption) {
	m.properties[key] = extractValueOptions(key, options...)
}

func WithDefaultValue(defaultValue string) ValueOption {
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
