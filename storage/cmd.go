package storage

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	. "recommend/util"
)

func (r RedisInstance) Exists(pat string) (res bool, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("keys %v", e)
		}
	}()
	res, err = redis.Bool(r.do("EXISTS", pat))

	return
}

func (r RedisInstance) Keys(pat string) (res []string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("keys %v", e)
		}
	}()
	res, err = redis.Strings(r.do("KEYS", pat))
	return
}

func (r RedisInstance) Mget(keys []string) (res []string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("keys %v", e)
		}
	}()

	newkeys := make([]interface{}, len(keys))
	for i, v := range keys {
		newkeys[i] = v
	}

	res, err = redis.Strings(r.do("MGET", newkeys...))

	return
}

func (r RedisInstance) get(pat string) (res string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("keys %v", e)
		}
	}()
	res, err = redis.String(r.do("GET", pat))
	return
}

func (r RedisInstance) Set(pat string, val string) (res string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("keys %v", e)
		}
	}()
	res, err = redis.String(r.do("SET", pat, val))
	return
}

func (r RedisInstance) Hmget(key string, fields []string) (res []string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	newfields := make([]interface{}, len(fields)+1)
	newfields[0] = key
	for i, v := range fields {
		newfields[i+1] = v
	}
	res, err = redis.Strings(r.do("HMGET", newfields...))
	Debugln("hmget\t", newfields, "\tres", res, "\terr", err)
	return
}

func (r RedisInstance) hgetAll(key string) (res map[string]string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	res, err = redis.StringMap(r.do("HGETALL", key))
	//Debugln("hgetAll", key, res, err)
	return
}

func (r RedisInstance) Lrange(key string, start, stop int) (res []string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	res, err = redis.Strings(r.do("lrange", key, start, stop))
	return
}

func (r RedisInstance) ZRANK(key string, field string) (res int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	res, err = redis.Int(r.do("zrank", key, field))
	return
}

func (r RedisInstance) ZSCORE(key string, field string) (res string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	res, err = redis.String(r.do("zscore", key, field))
	return
}

//zrange
func (r RedisInstance) Zrange(key string, min, max int, isWith bool) (res []string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	if isWith {
		res, err = redis.Strings(r.do("zrange", key, min, max, "WITHSCORES"))
	} else {
		res, err = redis.Strings(r.do("zrange", key, min, max))
	}
	return
}

func (r RedisInstance) ZADD(key string, fields []interface{}) (res int64, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	newfields := make([]interface{}, len(fields)+1)
	newfields[0] = key
	for i, v := range fields {
		newfields[i+1] = v
	}
	res, err = redis.Int64(r.do("ZADD", newfields...))
	Debugln("WZADD\t", newfields, "\tres", res, "\terr", err)
	return
}

func (r RedisInstance) LLEN(pat string) (res int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("keys %v", e)
		}
	}()
	res, err = redis.Int(r.do("LLEN", pat))
	return
}
func (r RedisInstance) RPOP(pat string) (res string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("keys %v", e)
		}
	}()
	res, err = redis.String(r.do("RPOP", pat))
	return
}

func (r RedisInstance) LPUSH(pat string, val string) (res string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("keys %v", e)
		}
	}()
	res, err = redis.String(r.do("LPUSH", pat, val))
	return
}

func (r RedisInstance) RPUSH(pat string, val string) (res string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("keys %v", e)
		}
	}()
	res, err = redis.String(r.do("RPUSH", pat, val))
	return
}

func (r RedisInstance) Hincrby(key string, fields []interface{}) (res int64, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	newfields := make([]interface{}, len(fields)+1)
	newfields[0] = key
	for i, v := range fields {
		newfields[i+1] = v
	}
	res, err = redis.Int64(r.do("HINCRBY", newfields...))
	Debugln("WHincrby\t", newfields, "\tres", res, "\terr", err)
	return
}

func (r RedisInstance) HMGET(key string, fields []string) (res []string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	newfields := make([]interface{}, len(fields)+1)
	newfields[0] = key
	for i, v := range fields {
		newfields[i+1] = v
	}
	res, err = redis.Strings(r.do("HMGET", newfields...))
	Debugln("WHMGET\t", newfields, "\tres", res, "\terr", err)
	return
}
func (r RedisInstance) HMSET(key string, fields []string) (res string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	newfields := make([]interface{}, len(fields)+1)
	newfields[0] = key
	for i, v := range fields {
		newfields[i+1] = v
	}
	res, err = redis.String(r.do("HMSET", newfields...))
	Debugln("WHMSET\t", newfields, "\tres", res, "\terr", err)
	return
}

func (r RedisInstance) Expire(key string, expire_sec int) (res bool, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	res, err = redis.Bool(r.do("expire", key, expire_sec))
	Debugln("WExpire\t", key, "\t", expire_sec, "\tres", res, "\terr", err)
	return
}

func (r RedisInstance) Eval(lua_script string, keycount int, fields []interface{}) (res interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	newfields := make([]interface{}, len(fields))
	for i, v := range fields {
		newfields[i] = v
	}

	res, err = r.doScript(lua_script, keycount, newfields...)
	Debugln("eval\t", lua_script, "\t", keycount, "\tres", res, "\terr", err)
	return
}

// eval "local ks = redis.call('keys',KEYS[1]) for i=1,#ks do redis.call('del',ks[i]) end return true" 1 mykey*
// eval "redis.call('hmset',KEYS[1],KEYS[2],KEYS[3]) redis.call('expire',KEYS[1]) return true" 3 key1 key2 key3
