package sample

import (
	"log/slog"

	"github.com/blueorb/config/config"
)

func DatabaseMap(log *slog.Logger) (cfgMap *config.ConfigMap) {
	cfgMap = config.NewConfigMap("database")
	cfgMap.Add("DB_CORE_CONNECTION_HOST", config.WithDefaultValue("localhost"))
	cfgMap.Add("DB_CORE_CREDENTIALS", config.WithSecretRef(SecretsDatabaseConnections, SecretsCoreDatabaseConnection))
	return
}
