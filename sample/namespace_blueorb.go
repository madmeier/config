package sample

import (
	"log/slog"

	"github.com/blueorb/config/config"
)

const (
	NameSpaceBlueOrb = "blueorb"
)

// NamespaceBlueOrb returns a new namespace configuration
func NamespaceBlueOrb(log *slog.Logger) *config.NameSpaceConfig {
	return config.NewNameSpaceConfig(NameSpaceBlueOrb)
}
