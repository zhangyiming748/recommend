package storage

import (
	. "recommend/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	key2redis map[string]RedisInstance
)

/**
  new Redis Client
*/
func NewRediss(runMode string) {
	key2redis = make(map[string]RedisInstance, 0)
	for i := 0; i < 10; i++ {
		redisName := "redis_" + fmt.Sprintf("%d", i)
		if GetVal(runMode+"_redis_args", redisName+"_args") != "" {
			newRedisHelper(runMode, redisName)
		}
	}
}

func newRedisHelper(runMode, redisName string) {
	section := runMode + "_redis_args"
	redisArgs := strings.Split(GetVal(section, redisName+"_args"), ";")
	redisUrl := GetVal(section, redisName+"_url")

	keysPrefix := strings.Split(strings.Split(redisArgs[0], ":")[1], ",")
	keysCacheable := strings.Split(strings.Split(redisArgs[1], ":")[1], ",")
	usePipe := strings.Split(redisArgs[2], ":")[1]

	prefixs := make([]Prefix, 0)
	for i, k := range keysPrefix {
		var p Prefix
		p.prefix = k
		t, _ := strconv.Atoi(keysCacheable[i])
		if t > 0 {
			p.isCacheable = true
			p.cacheMin, _ = time.ParseDuration(keysCacheable[i] + "m")
		}
		prefixs = append(prefixs, p)
	}
	r := RedisInit(prefixs, redisUrl, usePipe)
	for _, k := range keysPrefix {
		key2redis[k] = r
	}
	Infof(" connect to %s %s:  %s\r\n", runMode, redisName, redisUrl)
	return
}

func GetRedis(KeyPerfix string) RedisInstance {
	return key2redis[KeyPerfix]
}
