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

func mark(fab [][]int, sx, sy, dx, dy int) {
	for x := sx; x < sx + dx; x++ {
		for y := sy; y < sy + dy; y++ {
			fab[x][y]++
		}
	}
}


func main() {

	claims := make([]claim, 0)
	maxX, maxY := 0, 0

	// Read the claims and determine the fabric dimensions
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		c := strToClaim(scan.Text())
		claims = append(claims, c)
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
	fab := make([][]int, maxX)
	for x := 0; x < maxX; x++ {
		fab[x] = make([]int, maxY)
	}


	// Mark the claims
	for _, c := range claims {
		mark(fab, c.Left, c.Top, c.Width, c.Height)
	}

	// Count overlaps
	overlaps := 0
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if fab[x][y] > 1 {
				overlaps++
			}
		}
	}

	fmt.Println(overlaps)
}
