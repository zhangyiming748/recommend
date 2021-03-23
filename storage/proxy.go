package storage

import (
	. "../util"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"time"
)

type rCmd struct {
	cmd string
	key []interface{}
}
type rValue struct {
	val interface{}
	err error
}
type entry struct {
	rcmd      rCmd
	valueChan chan rValue
}

const (
	MaxCmds     = 100
	littleWhile = 5
	MaxLen      = 200
)

type Proxy struct {
	pool          *redis.Pool
	entryChan     []chan entry
	entryChanSize int
}

func NewProxy(pool *redis.Pool) *Proxy {
	return &Proxy{pool: pool, entryChan: make([]chan entry, 0)}
}

func (t *Proxy) proxy_init(size int) {
	Infoln("proxy_init, consumer chan len:", size)
	t.entryChanSize = size
	for i := 0; i < size; i++ {
		et := make(chan entry, MaxLen)
		t.entryChan = append(t.entryChan, et)
		go t.consumer(i)
	}
}

func (t *Proxy) consumer(i int) {
	robustconsumer := func() {
		defer PanicRecover(PanicPosition())
		cmdArray := make([]entry, 0)
		for {
			select {
			case rcmd := <-t.entryChan[i]:
				cmdArray = append(cmdArray, rcmd)
				if len(cmdArray) > MaxCmds {
					err := t.pipeline(cmdArray)
					if err != nil {
						Errorln("len(cmdArray):", len(cmdArray), "\terr:", err)
					}
					cmdArray = make([]entry, 0)
				}
			default:
				if len(cmdArray) == 0 {
					time.Sleep(littleWhile * time.Millisecond)
					continue
				} else {
					err := t.pipeline(cmdArray)
					if err != nil {
						Errorln("len(cmdArray):", len(cmdArray), "\terr:", err)
					}
					cmdArray = make([]entry, 0)
				}
			}
		}
	}
	for {
		robustconsumer()
		Errorln("should never be here...")
	}
}

func (t *Proxy) producer(cmd string, keys ...interface{}) (interface{}, error) {
	rc := rCmd{
		cmd: cmd,
		key: keys,
	}
	//Debugln("producer:", rc)
	ch := make(chan rValue)
	e := entry{
		rcmd:      rc,
		valueChan: ch,
	}
	t.entryChan[rand.Intn(t.entryChanSize)] <- e
	value := <-ch
	return value.val, value.err
}

func (t *Proxy) pipeline(entrys []entry) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	//Debugln(t.pool.ActiveCount())
	c := t.pool.Get()
	defer c.Close()
	//Debugln("debug pipeline size", len(entrys))
	for _, ent := range entrys {
		c.Send(ent.rcmd.cmd, ent.rcmd.key...)
		//Debugln("c.Send()", ent.rcmd.cmd, ent.rcmd.key)
	}
	c.Flush()
	for _, ent := range entrys {
		v, err := c.Receive()
		rv := rValue{
			val: v,
			err: err,
		}
		ent.valueChan <- rv
	}
	return
}
