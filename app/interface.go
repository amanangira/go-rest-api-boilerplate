package app

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type IAPI interface {
	IsDebug() bool
	GetLogger() *log.Logger
	GetDBClient() *sqlx.DB
}
