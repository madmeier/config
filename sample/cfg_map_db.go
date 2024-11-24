package sample

import (
	"log/slog"

	"github.com/blueorb/config/config"
)

func DatabaseMap(log *slog.Logger) (cfgMap *config.Map) {
	cfgMap = config.NewMap("database")
	cfgMap.Set("DB_CONNECTION_HOST", "blueorb")
	cfgMap.Set("DB_CREDENTIALS", "blueorb") // secret
	return
}
