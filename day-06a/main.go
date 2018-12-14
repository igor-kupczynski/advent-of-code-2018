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
	atPoint                 = 1
	distanceFromPoint       = 2
	equalDistanceFromPoints = 3
)

type status struct {
	distanceType int
	point        point
	value        int
}

func main() {
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

	// Build the board
	board := make([][]status, maxX)
	for x := 0; x < maxX; x++ {
		board[x] = make([]status, maxY)
		for y := 0; y < maxY; y++ {
			board[x][y] = status{
				distanceType: distanceFromPoint,
				point:        point{-1, -1},
				value:        99999,
			}
		}
	}

	// Calculate the distance for each point
	for _, p := range points {
		for x := 0; x < maxX; x++ {
			for y := 0; y < maxY; y++ {
				dist := intAbs(p.x-x) + intAbs(p.y-y)
				if dist == 0 {
					board[x][y].distanceType = atPoint
					board[x][y].point = p
					board[x][y].value = 0
				} else if dist < board[x][y].value {
					board[x][y].distanceType = distanceFromPoint
					board[x][y].point = p
					board[x][y].value = dist
				} else if dist == board[x][y].value {
					board[x][y].distanceType = equalDistanceFromPoints
					board[x][y].point = point{-1, -1}
				}
			}
		}
	}


	// Calculate the size for each area
	sizes := map[point]int{}
	infinites := map[point]bool{}

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if board[x][y].distanceType != equalDistanceFromPoints {
				p := board[x][y].point
				sizes[p]++
				if x == 0 || y == 0 || x + 1 == maxX || y + 1 == maxY {
					infinites[p] = true
				}
			}
		}
	}

	var maxAreaSize int
	var maxPoint point

	// Get largest, not infinite
	for point, areaSize := range sizes {
		if areaSize > maxAreaSize && !infinites[point] {
			maxPoint = point
			maxAreaSize = areaSize
		}
	}

	fmt.Printf("%v => %d\n", maxPoint, maxAreaSize)
}
