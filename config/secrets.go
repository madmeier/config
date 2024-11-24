package config

type Secret struct {
	SecretName string

	SecretKeys []string
}

func NewSecret(name string) *Secret {
	return &Secret{
		SecretName: name,
	}
}

func (s *Secret) Name() string {
	return s.SecretName
}

func (s *Secret) Add(key string) {
	s.SecretKeys = append(s.SecretKeys, key)
}
