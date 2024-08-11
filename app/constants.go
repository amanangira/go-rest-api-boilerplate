package app

const (
	// Environment placeholders

	EnvAWSRegion     = "REGION"
	EnvDebugKey      = "IS_DEBUG"
	EnvTrueString    = "true"
	EnvServerPortKey = "port"
	EnvStageKey      = "STAGE"
	EnvIsOffline     = "IS_OFFLINE"

	// PostgreSQL
	EnvDBSecretNameKey = "DB_SECRET_NAME"
	EnvDBHostKey       = "POSTGRESQL_HOST"
	EnvDBUsernameKey   = "POSTGRESQL_USERNAME"
	EnvDBPasswordKey   = "POSTGRESQL_PASSWORD"
	EnvDBPortKey       = "POSTGRESQL_PORT"
	EnvDBNameKey       = "POSTGRESQL_DB"
	EnvDBSSLMode       = "POSTGRESQL_SSL_MODE"
)
