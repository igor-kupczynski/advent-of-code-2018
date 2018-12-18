package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type rect struct {
	minX, maxX int
	minY, maxY int
}

func (r *rect) size() (x, y int) {
	x = r.maxX - r.minX
	y = r.maxY - r.minY
	return
}

func (r *rect) isLarger(other *rect) bool {
	ax, ay := r.size()
	bx, by := other.size()
	return ax >= bx && ay >= by && (ax > bx || ay > by)
}


type point struct {
	x, y   int
	vx, vy int
}

func (p *point) prev() {
	p.x -= p.vx
	p.y -= p.vy
}

func (p *point) next() {
	p.x += p.vx
	p.y += p.vy
}

type points struct {
	points     []*point
	size rect
}

func (p *points) move(f func(point *point)) bool {
	new := rect{}
	for _, point := range p.points {
		f(point)
		if point.x < new.minX {
			new.minX = point.x
		}
		if point.y < new.minY {
			new.minY = point.y
		}
		if point.x > new.maxX {
			new.maxX = point.x
		}
		if point.y > new.maxY {
			new.maxY = point.y
		}
	}
	prev := p.size
	p.size = new
	return new.isLarger(&prev)
}

func (p *points) next() bool {
	return p.move(func(point *point) {point.next()})
}

func (p *points) prev() bool {
	return p.move(func(point *point) {point.prev()})
}

func (p *points) canvas() [][]rune {
	dx, dy := -p.size.minX, -p.size.minY
	maxX, maxY := p.size.maxX+dx, p.size.maxY+dy
	canvas := make([][]rune, maxX+1)
	for x := 0; x <= maxX; x++ {
		canvas[x] = make([]rune, maxY+1)
		for y := 0; y <= maxY; y++ {
			canvas[x][y] = '.'
		}
	}
	for _, point := range p.points {
		canvas[point.x+dx][point.y+dy] = '#'
	}
	return canvas
}

func print(canvas [][]rune) {
	for x := 0; x < len(canvas); x++ {
		for y := 0; y < len(canvas[x]); y++ {
			fmt.Printf("%c", canvas[x][y])
		}
		fmt.Println()
	}
}

func main() {
	minX, maxX, minY, maxY := 0, 0, 0, 0
	items := []*point{}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var x, y, vx, vy int
		fmt.Sscanf(s.Text(), "position=<%d, %d> velocity=<%d, %d>", &x, &y, &vx, &vy)
		point := point{x, y, vx, vy}
		if point.x < minX {
			minX = point.x
		}
		if point.y < minY {
			minY = point.y
		}
		if point.x > maxX {
			maxX = point.x
		}
		if point.y > maxY {
			maxY = point.y
		}
		items = append(items, &point)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	points := points{items, rect{minX, maxX, minY, maxY}}

	for i := 0; i < 10000000; i++ {
		if (points.next()) {
			points.prev()
			print(points.canvas())
			fmt.Println("\nAfter", i, "seconds")
			return
		}
	}
}
