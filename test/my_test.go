package main

import "testing"

const (
	LevelDebug = (iota + 1) * 100
	LevelInfo
	LevelWarning
	LevelDatalog
	LevelError
	LevelReport
)

func TestMyFun(t *testing.T) {
	t.Logf("%v\n%v\n%v\n%v\n%v\n%v\n", LevelDebug, LevelInfo, LevelWarning, LevelDatalog, LevelError, LevelReport)
}
