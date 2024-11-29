package config

// default values per environment
var (
	envDefaultValues = make(map[string]map[string]any)
)

func AddEnvDefaultValue(env string, key string, value any) {
	if envDefaultValues[env] == nil {
		envDefaultValues[env] = make(map[string]any)
	}
	envDefaultValues[env][key] = value
}

func GetDefaultValue(env string, key string) (value any, ok bool) {
	envValues, ok := envDefaultValues[env]
	if !ok {
		return
	}
	value, ok = envValues[key]
	return
}
