package util

import "testing"

func TestGetVal(t *testing.T) {
	addr:=GetVal("proxy","address")
	t.Logf("address is %v",addr)
}
func TestFormat(t *testing.T) {

}