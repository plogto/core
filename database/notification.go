package database

import (
	"github.com/go-pg/pg/v10"
)

type Notification struct {
	DB *pg.DB
}
