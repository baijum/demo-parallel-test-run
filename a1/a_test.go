package a1

import "testing"

func TestHello(t *testing.T) {
	out := hello()
	if out != "hello" {
		t.Fail()
	}
}
