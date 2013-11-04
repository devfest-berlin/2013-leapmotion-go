package leapmotion

import (
	"testing"
)

func TestController(t *testing.T) {
	if _, err := NewController(); err != nil {
		t.Fatalf("Couldn't create a new Controller: %v\n", err)
	}
}
