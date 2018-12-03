package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var line string

	scanner := bufio.NewScanner(os.Stdin)
	point2 := 0
	point3 := 0
	for scanner.Scan() {
		line = scanner.Text()

		freqs := lineFreqs(line)
		p2, p3 := freqPoints(freqs)
		point2 = point2 + p2
		point3 = point3 + p3
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(point2 * point3)
}

func lineFreqs(line string) map[rune]int {
	freqs := make(map[rune]int)
	for _, v := range line {
		freq, ok := freqs[v]
		if ok {
			freqs[v] = freq + 1
		} else {
			freqs[v] = 1
		}
	}
	return freqs
}

func freqPoints(freqs map[rune]int) (p2, p3 int) {
	for _, v := range freqs {
		if v == 2 {
			p2 = 1
		} else if v == 3 {
			p3 = 1
		}
	}
	return
}
