package bytespool

import (
	"github.com/itsabgr/go-handy"
	"testing"
)

func TestBytesPool_Pull(t *testing.T) {
	for i := range handy.N(99) {
		max := i + 9
		b := Pull(uint(max))
		if len(b) > max {
			t.Fatalf("expected max len %d got %d", max, len(b))
		}
	}
}
