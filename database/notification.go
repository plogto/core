package database

import (
	"github.com/go-pg/pg"
)

type Notification struct {
	DB *pg.DB
}
