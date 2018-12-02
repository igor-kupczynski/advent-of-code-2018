package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	freq := 0
	seen := map[int]bool{}

	var line string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = scanner.Text()
		op := line[0]
		delta, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		if (op == '+') {
			freq = freq + delta
		} else if (op == '-') {
			freq = freq - delta
		}
		if _, ok := seen[freq]; ok == true {
			fmt.Println(freq)
			return
		}
		seen[freq] = true
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
