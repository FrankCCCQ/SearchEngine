package loading

import (
	"SearchEngine/config"
	log "SearchEngine/pkg/logger"
	"SearchEngine/repository/mysql/db"
	"SearchEngine/repository/redis"
)
func Loading() {
	config.InitConfig()
	log.InitLog()
	db.InitDB()
	redis.InitRedis()
	kfk.InitKafka()
}
