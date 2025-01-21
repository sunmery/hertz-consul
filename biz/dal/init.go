package dal

import (
	"hertz-consul/biz/dal/mysql"
	"hertz-consul/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
