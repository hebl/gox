package xweb

import (
	"fmt"

	"github.com/go-pg/pg"
)

// NewPg 新建一个PostgreSQL连接
func NewPg(dbconfig DBConfig) *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", dbconfig.Host, dbconfig.Port),
		Database: dbconfig.Database,
		User:     dbconfig.User,
		Password: dbconfig.Password,
		PoolSize: dbconfig.PoolSize,
	})
}
