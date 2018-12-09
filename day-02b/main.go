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
	scan := bufio.NewScanner(os.Stdin)
	var lines []string

	for scan.Scan() {
		line := scan.Text()
		lines = append(lines, line)
	}

	if err := scan.Err(); err != nil {
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
	if len(a) != len(b) {
		return math.MaxInt32
	}
	diff := 0
	for i := 0; i < len(a); i++ {
		if (a[i] != b[i]) {
			diff++
		}
	}
	return diff
}
