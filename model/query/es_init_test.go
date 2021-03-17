package query

import "testing"

func TestEsClient(t *testing.T) {
	node, _ :=EsClient()
	t.Log(node)
}