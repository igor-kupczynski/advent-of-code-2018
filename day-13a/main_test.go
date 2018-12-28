package main

import (
	"os"
	"strings"
	"testing"
)

func testBoard() *Board {
	var board *Board

	f, err := os.Open("example2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	board, err = ReadBoard(f)
	if err != nil {
		panic(err)
	}

	return board
}

func TestReadBoard(t *testing.T) {
	board := testBoard()

	// There should be two carts
	c1 := *board.carts[0]
	c2 := *board.carts[1]

	expected1 := Cart{0, 2, right, turnLeft}
	if c1 != expected1 {
		t.Errorf("Carts are not the same, expected %v, got %v.", expected1, c1)
	}

	expected2 := Cart{3, 9, down, turnLeft}
	if c2 != expected2 {
		t.Errorf("Carts are not the same, expected %v, got %v.", expected2, c2)
	}
}

func TestMoveCart(t *testing.T) {
	board := testBoard()

	// Let's focus on the first cart
	c := board.carts[0]

	path := []Cart{
		Cart{0, 3, right, turnLeft},
		Cart{0, 4, down, turnLeft},
		Cart{1, 4, down, turnLeft},
		Cart{2, 4, right, straight},
		Cart{2, 5, right, straight},
		Cart{2, 6, right, straight},
		Cart{2, 7, right, turnRight},
		Cart{2, 8, right, turnRight},
	}

	for _, expected := range path {
		board.moveCart(c)
		if *c != expected {
			builder := strings.Builder{}
			board.WriteBoard(&builder)
			t.Errorf("Invalid cart position, expected %v, got %v.\nBoard:\n%v\n\n\n",
				expected, c, builder.String())
		}
	}

}

func TestSortCarts(t *testing.T) {
	// d . b
	// c . .
	// . . a
	// expected order d, b, c, a
	var (
		a = Cart{x: 2, y: 2}
		b = Cart{x: 2, y: 0}
		c = Cart{x: 0, y: 1}
		d = Cart{x: 0, y: 0}
	)
	expected := []Cart{d, b, c, a}

	board := Board{carts: []*Cart{&a, &b, &c, &d}}

	board.sortCarts()
	actual := make([]Cart, 4)
	for i := 0; i < 4; i++ {
		actual[i] = *board.carts[i]
	}

	for i := 0; i < 4; i++ {
		if expected[i] != actual[i] {
			t.Errorf("Expected %v, but got %v instead (diff @ %d).\n", expected, actual, i)
		}
	}
}

func TestMove(t *testing.T) {
	board := testBoard()

	expectedMoves := 14
	var collision *Cart
	for i := 0; i < expectedMoves; i++ {
		collision = board.Move()
	}

	if collision == nil || collision.x != 7 || collision.y != 3 {
		builder := strings.Builder{}
		board.WriteBoard(&builder)
		t.Errorf("Invalid collision position, expected (7,3), got %v.\nBoard:\n%v\n\n\n",
			collision, builder.String())
	}
}
