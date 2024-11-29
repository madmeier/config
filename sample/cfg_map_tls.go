package sample

import (
	"log/slog"

	"github.com/blueorb/config/config"
)

const (
	CfgMapTLS = "tls"
)

func TLSMap(log *slog.Logger) (cfgMap *config.ConfigMap) {
	cfgMap = config.NewConfigMap(CfgMapTLS)
	cfgMap.Add("PROJECT_TLS_ACTIVE", config.WithDefaultValue(true))
	cfgMap.Add("PROJECT_TLS_CERT", config.WithDefaultValue("/app/secrets/star_domain.crt"))
	cfgMap.Add("PROJECT_TLS_KEY", config.WithDefaultValue("/app/secrets/star_domain_ch.key"))
	return
}
