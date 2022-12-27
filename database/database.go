package database

import (
	"context"
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
)

type DbLogger struct{}

func (d DbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d DbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	fq, _ := q.FormattedQuery()
	fmt.Println(string(fq))
	return nil
}

func New() *pg.DB {
	opt, err := pg.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	opt.PoolSize = 40
	return pg.Connect(opt)
}
