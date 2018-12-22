package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type gridID int

func hundreth(n int) int {
	return (n / 100) % 10
}

func (g gridID) powerLevel(x, y int) int {
	rackID := x + 10
	powerLevel := rackID * y
	powerLevel += int(g)
	powerLevel *= rackID
	powerLevel = hundreth(powerLevel) - 5
	return powerLevel
}

func regionPowerLevel(grid *[300][300]int, x, y, size int) int {
	sum := 0
	for dx := 0; dx < size; dx++ {
		for dy := 0; dy < size; dy++ {
			if (x+dx >= 300) || (y+dy >= 300) {
				continue
			}
			sum += grid[x+dx][y+dy]
		}
	}
	return sum
}

func calculateMaxRegion(grid *[300][300]int, size int) (int, int, int) {
	maxX, maxY, maxVal := 0, 0, math.MinInt32
	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			total := regionPowerLevel(grid, x, y, size)
			if total > maxVal {
				maxX = x + 1
				maxY = y + 1
				maxVal = total
			}
		}
	}
	return maxX, maxY, maxVal
}

func main() {
	var gid gridID
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		_ = s.Text()
		fmt.Sscanf(s.Text(), "%d", &gid)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	grid := [300][300]int{}
	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			grid[x-1][y-1] = gid.powerLevel(x, y)
		}
	}

	maxX, maxY, maxSize, maxTotal := 0, 0, 0, math.MinInt32
	for size := 1; size <= 300; size++ {
		tx, ty, tt := calculateMaxRegion(&grid, size)
		fmt.Printf("%03d (%d,%d) => %d | max:%d (%d,%d,%d)\n", size, tx, ty, tt, maxTotal, maxX, maxY, maxSize)
		if tt > maxTotal {
			maxX = tx
			maxY = ty
			maxTotal = tt
			maxSize = size
		}
	}
}
