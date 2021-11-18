package database

import (
	"os"

	"github.com/go-pg/pg/v10"
)

func New() *pg.DB {
	opt, err := pg.ParseURL(os.Getenv("DATABASE_URI"))
	if err != nil {
		panic(err)
	}

	return pg.Connect(opt)
}
