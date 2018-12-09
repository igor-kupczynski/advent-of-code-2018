package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	var nums []int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("Can't read %s: %v", s.Text(), err)
		}
		nums = append(nums, n)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	seen := map[int]bool{0: true}
	sum := 0
	for {
		for _, freq := range nums {
			sum += freq
			if _, isThere := seen[sum]; isThere == true {
				fmt.Println(sum)
				return
			}
			seen[sum] = true
		}
	}
}
