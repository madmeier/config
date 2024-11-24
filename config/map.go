package config

type Map struct {
	name string
	// environment variables with as generic default value
	properties map[string]string
}

func NewMap(name string) *Map {
	return &Map{
		name:       name,
		properties: map[string]string{},
	}
}

func (m *Map) Name() string {
	return m.name
}

func (m *Map) Properties() map[string]string {
	return m.properties
}

func (m *Map) Set(environmentVariable string, defaultValue string) {
	m.properties[environmentVariable] = defaultValue
}
