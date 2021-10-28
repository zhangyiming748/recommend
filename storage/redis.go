package storage

import (
	. "../util"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"strings"
	"time"
)

type Prefix struct {
	Prefix      string
	IsCacheable bool
	CacheMin    time.Duration
}

type RedisInstance struct {
	keyPrefix map[string]Prefix
	pool      *redis.Pool
	do        func(cmd string, keys ...interface{}) (interface{}, error)
	doScript  func(lua_script string, keycount int, keys ...interface{}) (interface{}, error)
	proxy     *Proxy
}

const (
	single_Client = 3
	more_Client   = 300
)

func RedisInit(prefixs []Prefix, redis_url, redis_usepipe string) RedisInstance {
	var r RedisInstance

	r.keyPrefix = make(map[string]Prefix)
	for _, p := range prefixs {
		r.keyPrefix[p.Prefix] = p
	}

	redis_connects := strings.Split(redis_url, "?")
	redis_pass := strings.Split(redis_connects[1], "&")[0]
	redis_pass = strings.Split(redis_pass, "=")[1]
	redis_db, _ := strconv.Atoi(strings.Split(strings.Split(redis_connects[1], "&")[1], "=")[1])
	redis_connect_str := redis_connects[0]
	if strings.ToLower(redis_usepipe) == "true" {
		r.pool = newPool(redis_connect_str, redis_pass, redis_db, single_Client)
		r.proxy = NewProxy(r.pool)
		r.do = r.proxy.producer
		r.proxy.proxy_init(single_Client)
	} else {
		r.pool = newPool(redis_connect_str, redis_pass, redis_db, more_Client)
		r.do = r.redisDo
	}
	r.doScript = r.redisScript
	return r
}

func newPool(redis_connect_str, redis_pass string, redis_db, clients int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     clients,
		Wait:        true,
		IdleTimeout: 240 * time.Second,
		MaxActive:   clients, // max number of connections
		Dial: func() (redis.Conn, error) {
			db, passwd, connTimeout, readTimeout := redis.DialDatabase(redis_db), redis.DialPassword(redis_pass), redis.DialConnectTimeout(600*time.Second), redis.DialReadTimeout(60*time.Second)
			c, err := redis.Dial("tcp", redis_connect_str, db, passwd, connTimeout, readTimeout)
			if err != nil {
				Errorln("redispool dial err:", err)
			}
			return c, err
		},
		//Use the TestOnBorrow function to check the health of an idle connection before the connection is returned to the application.
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func (r RedisInstance) redisScript(src string, keycount int, keys ...interface{}) (v interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("redisScript %v", e)
		}
	}()
	c := r.pool.Get()
	defer c.Close()
	var getScript = redis.NewScript(keycount, src)
	return getScript.Do(c, keys...)
}

func (r RedisInstance) redisDo(cmd string, keys ...interface{}) (v interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("redisDo %v", e)
		}
	}()
	c := r.pool.Get()
	defer c.Close()
	return c.Do(cmd, keys...)
}

func (r RedisInstance) isCacheable(keyPrefix string) (bool, time.Duration) {
	p, ok := r.keyPrefix[keyPrefix]
	if ok {
		return p.IsCacheable, p.CacheMin
	}
	return false, 0
}
