package sample

import (
	"log/slog"

	"github.com/blueorb/config/config"
)

func GoogleProjectMap(log *slog.Logger) (cfgMap *config.ConfigMap) {
	cfgMap = config.NewConfigMap("google")
	cfgMap.Add("GOOGLE_PROJECT_ID", config.WithDefaultValue("blueorb"))
	cfgMap.Add("GOOGLE_PROJECT_NAME", config.WithDefaultValue("blueorb.json"))
	return
}
