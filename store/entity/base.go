package entity

import (
	"database/sql"
	"time"
)

type Base struct {
	ID        string       `db:"id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
