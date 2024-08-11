package app

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// TODO - setup sqlx.DB client
func GetDBClient() *sqlx.DB {
	return &sqlx.DB{}
}

func DefaultDBDSN() string {
	port := os.Getenv(EnvDBPortKey)
	dbName := os.Getenv(EnvDBNameKey)
	hostname := os.Getenv(EnvDBHostKey)
	username := os.Getenv(EnvDBUsernameKey)
	password := os.Getenv(EnvDBPasswordKey)

	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s", username, dbName, password, hostname, port)

	if sslMode := os.Getenv(EnvDBSSLMode); sslMode != "" {
		dsn += " sslmode=" + sslMode
	}

	return dsn
}

func NewSqlxClient() *sqlx.DB {
	return sqlx.MustConnect("postgres", DefaultDBDSN())
}
