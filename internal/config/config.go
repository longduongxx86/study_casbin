package config

import "github.com/zeromicro/go-zero/rest"

type SqlDB struct {
	SQL_USER     string
	SQL_PASSWORD string
	SQL_PORT     string
	SQL_DBNAME   string
}

type Rdb struct {
	RDB_HOST string
	RDB_PASS string
	RDB_TYPE string
}

type Config struct {
	Rest  rest.RestConf
	SqlDB SqlDB
	Rdb   Rdb
	
	Auth  struct {
		AccessSecret string
		AccessExpire int64
	}
}
