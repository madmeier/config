package sample

import (
	"log/slog"

	"github.com/blueorb/config/config"
)

const (
	SecretsDatabaseConnections    = "database_connection_strings"
	SecretsCoreDatabaseConnection = "core_db_connection"
)

func SecretsDatabase(log *slog.Logger) (secret *config.Secret) {
	secret = config.NewSecret(SecretsDatabaseConnections)
	secret.Add(SecretsCoreDatabaseConnection)
	return
}
