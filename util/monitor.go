package util

import (
	"encoding/json"
	"fmt"
	"sync/atomic"
	"time"
)

/***
  main 方法中引用
*/
const (
	LatencySize = 1000
	QpsDuration = 600
)

type AtomicUint64 uint64

func (i *AtomicUint64) Add(delta uint64) uint64 {
	return atomic.AddUint64((*uint64)(i), delta)
}

func (i *AtomicUint64) Set(val uint64) {
	atomic.StoreUint64((*uint64)(i), val)
}

func (i *AtomicUint64) Load() uint64 {
	return atomic.LoadUint64((*uint64)(i))
}

func avgLatency(latencies [LatencySize]time.Duration) time.Duration {
	var d time.Duration = 0
	for _, v := range latencies {
		d += v
	}
	return d / LatencySize
}

type Server struct {
	Hits         AtomicUint64
	Qps          AtomicUint64
	fails        AtomicUint64
	latencyIndex AtomicUint64
	latencies    [LatencySize]time.Duration
	Start        string
}

func NewMonitorServer() Server {

	server := Server{
		Start: time.Now().Format("2006-01-02 15:04:05"),
	}
	go server.Cycle()

	return server
}
//设置延迟函数
func (this *Server) SetLatency(duration time.Duration) {
	pos := this.latencyIndex.Add(1)
	this.latencies[pos%LatencySize] = duration
}
//自定义解码json
func (this Server) MarshalJSON() ([]byte, error) {
	hits := this.Hits.Load()
	since := time.Now().Unix() % QpsDuration
	qps := float64(this.Qps.Load()) / float64(since)

	index := this.latencyIndex.Load()

	lastTen := make([]string, 10)
	var i uint64 = 0
	for {
		lastTen[i] = fmt.Sprintf("%v", this.latencies[(index-10+i)%LatencySize])
		i += 1
		if i == 10 {
			break
		}
	}

	return json.Marshal(map[string]interface{}{
		"ServerStart":    this.Start,
		"Hits":           hits,
		"Qps":            fmt.Sprintf("%.2f", qps),
		"LastTenLatency": lastTen,
		"avgLatency":     fmt.Sprintf("%v", avgLatency(this.latencies)),
	})
}
//查看qps(每秒查询率)
func (this *Server) Cycle() {
	for {
		now := time.Now().Unix()
		next := (now/QpsDuration)*QpsDuration + QpsDuration
		<-time.After(time.Duration(next-now) * time.Second)
		this.Qps.Set(0)
	}
}
