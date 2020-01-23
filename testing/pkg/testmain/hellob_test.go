package testmain

import "testing"

func TestHelloB(t *testing.T) {
	if HelloB() != "HelloB" {
		t.Fatal("error")
	}
}
