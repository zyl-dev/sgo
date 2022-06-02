package dbs

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

// 数据库的初始化
// 依赖于 gokit (https://github.com/btccom/gokit)
var (
	// DBMaps 代表 sqlx 连接的 map
	DBMaps map[string]*sqlx.DB

	// RedisDB redis client
	RedisDB        *redis.Client
	RedisClusterDB *redis.ClusterClient
	RedisRingDB    *redis.Ring
)

// CheckDBConnExists 检查连接是否存在
func CheckDBConnExists(conn string) bool {
	_, ok := DBMaps[conn]
	return ok
}
