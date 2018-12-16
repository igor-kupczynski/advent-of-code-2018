package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

type point struct {
	x, y int
}

func text2point(line string) point {
	var x, y int
	fmt.Sscanf(line, "%d, %d", &x, &y)
	return point{x, y}
}

func intAbs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

const (
	allowedDist = 10000
)

func main() {
	var regionSize int
	var maxX, maxY int
	points := []point{}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		p := text2point(s.Text())
		points = append(points, p)
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}


	// We need to account for the "border" area
	maxX++
	maxY++

	// Iterate over the board
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			allPointsDist := 0
			for _, p := range points {
				allPointsDist += intAbs(p.x - x) + intAbs(p.y - y)
				if allPointsDist >= allowedDist {
					break
				}
			}
			if allPointsDist < allowedDist {
				regionSize++
			}
		}
	}

	fmt.Println(regionSize)
}
