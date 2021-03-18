package recommend

var MaxTraffic uint64
var TrafficCtrl AtomicUint64

func trafficBegin() {
	TrafficCtrl.Add(1)
}

func trafficDone() {
	var i int64 = -1
	TrafficCtrl.Add(uint64(int64(i)))
}

func isTrafficJam() bool {
	//	Errorln(TrafficCtrl.Load())
	return TrafficCtrl.Load() > MaxTraffic
}

func SetMaxTraffic(maxi uint64) {
	//限制的最大并发流量, 超出则降级
	MaxTraffic = maxi
}

func init() {
	TrafficCtrl.Set(uint64(8)) //for safe, 防止uint回卷
}
