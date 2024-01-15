package database

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DBConn struct {
	Conn      sqlx.SqlConn
	ConnCache sqlc.CachedConn
}

func Connect(dataSource string, conf cache.CacheConf) *DBConn {
	sqlConn := sqlx.NewMysql(dataSource)
	d := &DBConn{
		Conn: sqlConn,
	}
	if conf != nil {
		cachedConn := sqlc.NewConn(sqlConn, conf)
		d.ConnCache = cachedConn
	}
	return d
}
