package util

import "testing"

func TestGetVal(t *testing.T) {
	s := GetVal("dev_redis_args", "redis_0_args")
	t.Log(s)
}
