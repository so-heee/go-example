package testmain

import "testing"

func TestHelloA(t *testing.T) {
	if HelloA() != "HelloA" {
		t.Fatal("error")
	}
}
