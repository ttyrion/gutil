package maputil

import (
	"testing"
)

func TestReadMapValue(t *testing.T) {
	m := map[string]interface{} {
		"a": "hello",
	}
	v,_ := ReadMapValue(m, "a")
	if v != "hello" {
		t.Fatalf("ReadMapValue failed.")
	}
}