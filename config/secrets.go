package config

import "log/slog"

type Secret struct {
	secretName string
	nameSpace  string
	secretKeys []string
}

func NewSecret(name string, nameSpace ...string) *Secret {
	return &Secret{
		nameSpace:  getNameSpace(nameSpace...),
		secretName: name,
	}
}

func (s *Secret) Name() string {
	return s.secretName
}

func (s *Secret) NameSpace() string {
	return s.nameSpace
}

func (s *Secret) Add(key string) {
	s.secretKeys = append(s.secretKeys, key)
}

func (a *ArchConfig) CheckSecrets(log *slog.Logger) {
}
