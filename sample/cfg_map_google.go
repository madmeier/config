package sample

import (
	"log/slog"

	"github.com/blueorb/config/config"
)

func GoogleProjectMap(log *slog.Logger) (cfgMap *config.Map) {
	cfgMap = config.NewMap("google")
	cfgMap.Add("GOOGLE_PROJECT_ID", config.WithDefaultValue("blueorb"))
	cfgMap.Add("GOOGLE_PROJECT_NAME", config.WithDefaultValue("blueorb.json"))
	return
}
