package bpm

import (
	"fmt"
	"testing"
)

func TestHasKey(t *testing.T) {
	list := newCircularList(10)

	list.insert(1, true)
	list.insert(2, true)
	list.insert(4, true)

	var tests = []struct {
		a    int
		b    bool
		want bool
	}{
		{1, true, true},
		{2, true, true},
		{3, true, false},
		{4, true, true},
		{5, true, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := list.hasKey(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestSize(t *testing.T) {
	list := newCircularList(10)

	ans := list.size
	if ans != 0 {
		t.Errorf("got %d, want %d", ans, 0)
	}

	list.insert(1, true)
	list.insert(1, true)
	list.insert(1, true)
	list.insert(1, true)
	list.insert(2, true)
	list.insert(4, true)

	ans = list.size
	if ans != 3 {
		t.Errorf("got %d, want %d", ans, 3)
	}
}
