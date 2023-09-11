package svc

import (
	"docker_gozero/internal/config"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config  config.Config
	MysqlDB sqlx.SqlConn
	RedisDB redis.Redis
	UserTokenMiddleware rest.Middleware
}

func connDB(sct *config.Config) sqlx.SqlConn {
	sql := sct.SqlDB
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		sql.SQL_USER, sql.SQL_PASSWORD, sql.SQL_PORT, sql.SQL_DBNAME,
	)
	conn := sqlx.NewMysql(dsn)
	return conn
}

func connRDB(sct *config.Config) redis.Redis {
	conf := redis.RedisConf{
		Host:        "127.0.0.1:6379",
		Pass:        "",
		// Host: sct.Rdb.RDB_HOST,
		// Pass: sct.Rdb.RDB_PASS,

		Type: sct.Rdb.RDB_TYPE,
	}
	rdb := redis.MustNewRedis(conf)
	return *rdb
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		MysqlDB: connDB(&c),
		RedisDB: connRDB(&c),
	}
}
