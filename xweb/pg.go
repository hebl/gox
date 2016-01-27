package xweb

import "gopkg.in/pg.v3"

// NewPg 新建一个PostgreSQL连接
func NewPg(dbconfig DBConfig) *pg.DB {
	return pg.Connect(&pg.Options{
		Host:     dbconfig.Host,
		Port:     dbconfig.Port,
		Database: dbconfig.Database,
		User:     dbconfig.User,
		Password: dbconfig.Password,
		PoolSize: dbconfig.PoolSize,
	})
}
