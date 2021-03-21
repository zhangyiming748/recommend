package main

import "testing"

func TestEsClient(t *testing.T) {
	node, _ := EsClient()
	t.Log(node)
}
func TestQFromeES(t *testing.T) {
	qFromES()
}