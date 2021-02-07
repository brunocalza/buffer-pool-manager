package bpm

import (
	"testing"
)

func TestSample(t *testing.T) {
	clockReplacer := NewClockReplacer(10)
	clockReplacer.Unpin(1)
	clockReplacer.Unpin(2)
	clockReplacer.Unpin(3)
	clockReplacer.Unpin(4)
	clockReplacer.Unpin(5)
	clockReplacer.Unpin(6)
	clockReplacer.Unpin(1)

	ans := clockReplacer.Size()
	if ans != 6 {
		t.Errorf("got %d, want %d", ans, 6)
	}

	val := clockReplacer.Victim()
	if *val != 1 {
		t.Errorf("got %d, want %d", *val, 1)
	}

	val = clockReplacer.Victim()
	if *val != 2 {
		t.Errorf("got %d, want %d", *val, 2)
	}

	val = clockReplacer.Victim()
	if *val != 3 {
		t.Errorf("got %d, want %d", *val, 3)
	}

	clockReplacer.Pin(3)
	clockReplacer.Pin(4)
	ans = clockReplacer.Size()
	if ans != 2 {
		t.Errorf("got %d, want %d", ans, 2)
	}

	clockReplacer.Unpin(4)

	val = clockReplacer.Victim()
	if *val != 5 {
		t.Errorf("got %d, want %d", *val, 5)
	}

	val = clockReplacer.Victim()
	if *val != 6 {
		t.Errorf("got %d, want %d", *val, 6)
	}

	val = clockReplacer.Victim()
	if *val != 4 {
		t.Errorf("got %d, want %d", *val, 4)
	}
}
