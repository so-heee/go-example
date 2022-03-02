package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	if Hello() != "Hello" {
		t.Fatal("error")
	}
}
