package main

import (
	"testing"
)

func TestHundreth(t *testing.T) {
	tables := []struct {
		in, out int
	}{
		{12345, 3},
		{12, 0},
		{949, 9},
	}
	for _, tab := range tables {
		if h := hundreth(tab.in); h != tab.out {
			t.Errorf("hundreth(%d) should be %d, but is %d\n", tab.in, tab.out, h)
		}
	}
}

func TestPowerLevel(t *testing.T) {
	tables := []struct {
		gid  gridID
		x, y int
		out  int
	}{
		{8, 3, 5, 4},
		{57, 122, 79, -5},
		{39, 217, 196, 0},
		{71, 101, 153, 4},
	}
	for _, tab := range tables {
		if level := tab.gid.powerLevel(tab.x, tab.y); level != tab.out {
			t.Errorf("gid=%d powerlevel@%d,%d should be %d, but is %d\n", tab.gid, tab.x, tab.y, tab.out, level)
		}
	}
}
