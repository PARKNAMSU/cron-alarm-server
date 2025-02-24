package common_entity

import (
	"database/sql"
	"time"
)

type Timestamp struct {
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
