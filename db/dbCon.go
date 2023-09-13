package db

import (
	"github.com/aminkhn/portfolio/db/cache"
	"github.com/aminkhn/portfolio/db/sql"
)

func SetupDatabase() {
	cache.RedisConnet()
	sql.MysqlConnect()
}
func DisconnectDatabase() {
	cache.RedisClose()
	sql.MysqlClose()
}
