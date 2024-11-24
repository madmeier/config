package sample

import (
	"log/slog"

	"github.com/blueorb/config/config"
)

func GoogleProjectMap(log *slog.Logger) (cfgMap *config.Map) {
	cfgMap = config.NewMap("google")
	cfgMap.Set("GOOGLE_PROJECT_ID", "blueorb")
	cfgMap.Set("GOOGLE_APPLICATION_CREDENTIALS", "blueorb") // secret
	return
}
