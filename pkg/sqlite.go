package pkg

import (
	"database/sql"
)

type Sqlite interface {
	Client() *sql.DB
}
