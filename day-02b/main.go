package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"math"
)

func main() {
	lines := readLines()

	a, b := findDist1(lines)


	fmt.Println(a)
	fmt.Println(b)
}

func readLines() []string {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0, 1000)

	var idx = 0
	for scanner.Scan() {
		lines = lines[0:idx+1]
		lines[idx] = scanner.Text()
		idx++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func findDist1(lines []string) (lineA, lineB string) {
	for a := 0; a < len(lines); a++ {
		for b := a + 1; b < len(lines); b++ {
			lineA = lines[a]
			lineB = lines[b]
			dist := distance(lineA, lineB)
			if dist == 1 {
				return
			}
		}
	}
	return
}


func distance(a string, b string) int {
	diff := 0
	for i := 0; i < int(math.Min(float64(len(a)), float64(len(b)))); i++ {
		if (a[i] != b[i]) {
			diff++
		}
	}
	diff = diff + int(math.Abs(float64(len(a)) - float64(len(b))))
	return diff
}
