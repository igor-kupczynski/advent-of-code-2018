package main

import (
	"os"
	"bufio"
	"log"
	"fmt"
)

func main() {
	sum := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatal(err)
		}
		sum += n
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
