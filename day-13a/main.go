package main

import (
	"bufio"
	"io"
	"fmt"
	"os"
	"log"
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


// Cart represents a cart moving on the track
type Cart struct {
	y, x int
	direction byte
}

// Board depicts the track layout on which the carts move
type Board struct {
	tracks [][]byte
	width, height int
	carts []*Cart
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
		buf[c.y][c.x] = c.direction
	}

	// Copy carts
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

func main() {
	b, err := ReadBoard(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	if err := b.WriteBoard(os.Stdout); err != nil {
		log.Fatalln(err)
	}
}
