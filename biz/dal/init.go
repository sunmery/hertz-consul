package dal

import (
	"github.com/sunmery/hertz-consul/biz/dal/mysql"
	"github.com/sunmery/hertz-consul/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
