package storage

import (
	"github.com/pmylund/go-cache"
	"strings"
	"time"
)

const (
	DEFAULTEXPIRATION = 3 * time.Minute //缓存失效时间
	PURGESEXPIRED     = 15 * time.Minute
	MAXEXPIRATION     = 360 * time.Minute
)

var (
	c *cache.Cache = cache.New(DEFAULTEXPIRATION, PURGESEXPIRED)
)

/**
  缓存内容
*/
func FetchCacheContent(key string) (val interface{}, f bool) {
	v, f := c.Get(key)
	if f {
		val = v
	}
	return
}
func SetCacheContent(key string, content interface{}, d time.Duration) {
	c.Add(key, content, d)
}

func (r RedisInstance) HgetAll(key string) (val map[string]string, err error) {
	prefix := strings.Split(key, ":")[0]
	isCache, cacheMin := r.isCacheable(prefix)
	if isCache {
		val_, found := c.Get(key)
		if !found {
			val, err = r.hgetAll(key)
			if err == nil && len(val) > 0 {
				c.Add(key, val, cacheMin)
			}
		} else {
			val = val_.(map[string]string)
		}
		return
	} else {
		return r.hgetAll(key)
	}
	return
}

func (r RedisInstance) Get(key string) (val string, err error) {
	prefix := strings.Split(key, ":")[0]
	isCache, cacheMin := r.isCacheable(prefix)
	if isCache {
		val_, found := c.Get(key)
		if !found {
			val, err = r.get(key)
			if err == nil && len(val) > 0 {
				c.Add(key, val, cacheMin)
			}
		} else {
			val = val_.(string)
		}
		return
	} else {
		return r.get(key)
	}
	return
}
