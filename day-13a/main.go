package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

// Tracks
const (
	empty = ' '
	cross = '+'

	// From north
	ne = '\\'
	ns = '|'
	nw = '/'

	// From east
	en = ne
	ew = '-'
	es = nw

	// From south
	se = nw
	sn = ns
	sw = ne

	// From west
	wn = nw
	we = ew
	ws = ne
)

// Carts
const (
	up    = '^'
	down  = 'v'
	left  = '<'
	right = '>'
)

// Turns
const (
	turnLeft  = 0
	straight  = 1
	turnRight = 2
)

// Cart represents a cart moving on the track
type Cart struct {
	y, x      int
	direction byte
	nextTurn  byte
}

func (c Cart) String() string {
	var turn string

	switch c.nextTurn {
	case turnLeft:
		turn = "left"
	case straight:
		turn = "straight"
	case turnRight:
		turn = "right"
	}

	return fmt.Sprintf("{%d,%d %c %v}", c.x, c.y, c.direction, turn)
}

// Board depicts the track layout on which the carts move
type Board struct {
	tracks        [][]byte
	width, height int
	carts         []*Cart
}

// ReadBoard reads the board from a reader
func ReadBoard(r io.Reader) (*Board, error) {
	tracks := [][]byte{}
	carts := []*Cart{}

	s := bufio.NewScanner(r)

	width, height := 0, 0
	for s.Scan() {
		line := s.Text()
		if len(line) > width {
			width = len(line)
		}
		tracks = append(tracks, make([]byte, width))
		for x := 0; x < width; x++ {
			switch ch := line[x]; ch {
			case up, down:
				tracks[height][x] = ns
				carts = append(carts, &Cart{y: height, x: x, direction: ch})
			case left, right:
				tracks[height][x] = ew
				carts = append(carts, &Cart{y: height, x: x, direction: ch})
			case empty, cross, ne, ns, nw, ew:
				tracks[height][x] = ch
			default:
				panic(fmt.Sprintf("Illegal track component '%c'\n", ch))
			}
		}
		height++
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	b := &Board{tracks, width, height, carts}
	return b, nil
}

// WriteBoard writes the board representation to a writer
func (b *Board) WriteBoard(w io.Writer) error {
	buf := make([][]byte, b.height)

	// Copy tracks
	for y := 0; y < b.height; y++ {
		buf[y] = make([]byte, b.width)
		for x, ch := range b.tracks[y] {
			buf[y][x] = ch
		}
	}

	// Copy carts
	for _, c := range b.carts {
		switch buf[c.y][c.x] {
		case up, right, down, left:
			// collision
			buf[c.y][c.x] = 'X'
		default:
			buf[c.y][c.x] = c.direction
		}
	}

	for y := 0; y < b.height; y++ {
		if _, err := w.Write(buf[y]); err != nil {
			return err
		}
		if _, err := w.Write([]byte{'\n'}); err != nil {
			return err
		}
	}
	return nil
}

// Move moves all the carts on the board
//  it stops in case of a collision and returns the offending cart
func (b *Board) Move() *Cart {
	b.sortCarts()
	for i, c := range b.carts {
		b.moveCart(c)
		// Collision?
		for j, other := range b.carts {
			if i == j {
				continue
			}
			if c.x == other.x && c.y == other.y {
				return c
			}
		}
	}
	return nil
}

func (b *Board) moveCart(c *Cart) {
	// move and turn according to tracks
	switch c.direction {
	case left:
		c.x--
		switch b.tracks[c.y][c.x] {
		case en:
			c.direction = up
		case es:
			c.direction = down
		}
	case right:
		c.x++
		switch b.tracks[c.y][c.x] {
		case wn:
			c.direction = up
		case ws:
			c.direction = down
		}
	case up:
		c.y--
		switch b.tracks[c.y][c.x] {
		case sw:
			c.direction = left
		case se:
			c.direction = right
		}
	case down:
		c.y++
		switch b.tracks[c.y][c.x] {
		case nw:
			c.direction = left
		case ne:
			c.direction = right
		}
	}

	// maybe turn on intersection
	if b.tracks[c.y][c.x] == cross {
		switch c.direction {
		case up:
			switch c.nextTurn {
			case turnLeft:
				c.direction = left
			case turnRight:
				c.direction = right
			}
		case right:
			switch c.nextTurn {
			case turnLeft:
				c.direction = up
			case turnRight:
				c.direction = down
			}
		case down:
			switch c.nextTurn {
			case turnLeft:
				c.direction = right
			case turnRight:
				c.direction = left
			}
		case left:
			switch c.nextTurn {
			case turnLeft:
				c.direction = down
			case turnRight:
				c.direction = up
			}
		}
		c.nextTurn = (c.nextTurn + 1) % 3
	}
}

func (b *Board) sortCarts() {
	sort.Slice(b.carts, func(i, j int) bool {
		if b.carts[i].y < b.carts[j].y {
			return true
		}
		if b.carts[i].y > b.carts[j].y {
			return false
		}
		return b.carts[i].x < b.carts[j].x
	})
}

func main() {
	board, err := ReadBoard(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	var collision *Cart
	for collision == nil {
		collision = board.Move()
	}

	builder := strings.Builder{}
	board.WriteBoard(&builder)
	fmt.Printf("Board:\n%v\n\nSolution: %d,%d\n", builder.String(), collision.x, collision.y)
}
