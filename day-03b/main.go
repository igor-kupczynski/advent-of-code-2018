package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
)

type claim struct {
	Num int
	Left, Top int
	Width, Height int
}

func (c *claim) maxX() int {
	return c.Left + c.Width
}

func (c *claim) maxY() int {
	return c.Top + c.Height
}

func strToClaim(text string) claim {
	c := claim{}
	_, err := fmt.Sscanf(text, "#%d @ %d,%d: %dx%d", &c.Num, &c.Left, &c.Top, &c.Width, &c.Height)
	if err != nil {
		log.Fatalf("Can't parse %s: %v", text, err)
	}
	return c
}

func mark(fab [][][]claim, c *claim) {
	for x := c.Left; x < c.Left + c.Width; x++ {
		for y := c.Top; y < c.Top + c.Height; y++ {
			fab[x][y] = append(fab[x][y], *c)
		}
	}
}


func main() {

	claimOverlaps := make(map[claim]bool)
	maxX, maxY := 0, 0

	// Read the claims and determine the fabric dimensions
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		c := strToClaim(scan.Text())
		claimOverlaps[c] = false
		if c.maxX() > maxX {
			maxX = c.maxX()
		}
		if c.maxY() > maxY {
			maxY = c.maxY()
		}
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	// Make the fabric
	fab := make([][][]claim, maxX)
	for x := 0; x < maxX; x++ {
		fab[x] = make([][]claim, maxY)
		for y := 0; y < maxY; y++ {
			fab[x][y] = make([]claim, 0)
		}
	}


	// Mark the claims
	for c := range claimOverlaps {
		mark(fab, &c)
	}

	// Count overlaps
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if point := fab[x][y]; len(point) > 1 {
				for _, c := range point {
					claimOverlaps[c] = true
				}
			}
		}
	}

	for k, v := range claimOverlaps {
		if v == false {
			fmt.Println(k.Num)
		}
	}
}
