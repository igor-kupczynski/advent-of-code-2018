package main

import (
	"bufio"
	"os"
	"log"
	"fmt"
	"math"
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


func regionPowerLevel(grid [300][300]int, x, y int) int {
	sum := 0
	for dx := 0; dx < 3; dx++ {
		for dy := 0; dy < 3; dy++ {
			if (x + dx >= 300) || (y +dy >= 300) {
				continue
			}
			sum += grid[x+dx][y+dy]
		}
	}
	return sum
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


	maxX, maxY, maxVal := 0, 0, math.MinInt32
	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			total := regionPowerLevel(grid, x, y)
			if total > maxVal {
				maxX = x + 1
				maxY = y + 1
				maxVal = total
			}
		}
	}

	fmt.Println(maxX, maxY, maxVal)
}
