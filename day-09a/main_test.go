package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	// We start with a single marble
	game := makeCircle()
	expt := []int{0}
	if !reflect.DeepEqual(game.Describe(), expt) {
		t.Error(game.Describe(), expt)
	}

	// Let's add a marble
	game.Add(makeMarble(1))
	expt = []int{1, 0}
	if !reflect.DeepEqual(game.Describe(), expt) {
		t.Error(game.Describe(), expt)
	}

	// And some more
	game.Add(makeMarble(2))
	game.Add(makeMarble(3))
	game.Add(makeMarble(4))
	game.Add(makeMarble(5))
	game.Add(makeMarble(6))
	expt = []int{6, 3, 0, 4, 2, 5, 1}
	if !reflect.DeepEqual(game.Describe(), expt) {
		t.Error(game.Describe(), expt)
	}
}

func TestRemove(t *testing.T) {
	// Let's add marbles up to 22
	game := makeCircle()
	for marble := 1; marble < 23; marble++ {
		game.Add(makeMarble(marble))
	}

	removed := game.Remove()
	if removed != 9 {
		t.Error(removed, 9)
	}

	expt := []int{19, 2, 20, 10, 21, 5, 22, 11, 1, 12, 6, 13, 3, 14, 7, 15, 0, 16, 8, 17, 4, 18}
	if !reflect.DeepEqual(game.Describe(), expt) {
		t.Error(game.Describe(), expt)
	}
}

func TestExamples(t *testing.T) {
	tables := []struct {
		players   int
		last      int
		highScore int64
	}{
		{10, 1618, int64(8317)},
		{13, 7999, int64(146373)},
		{17, 1104, int64(2764)},
		{21, 6111, int64(54718)},
		{30, 5807, int64(37305)},
	}

	for _, tab := range tables {
		act := play(tab.players, tab.last)
		if act != tab.highScore {
			t.Errorf("Game %v, got %v", tab, act)
		}
	}
}
