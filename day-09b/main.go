package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type marble struct {
	next  *marble
	prev  *marble
	value int
}

func makeMarble(val int) *marble {
	return &marble{value: val}
}

type circle struct {
	head *marble
}

func makeCircle() *circle {
	m := makeMarble(0)
	m.next = m
	m.prev = m

	return &circle{m}
}

func (c *circle) Describe() []int {
	result := []int{}
	curr := c.head
	for {
		result = append(result, curr.value)
		curr = curr.next
		if curr == c.head {
			return result
		}
	}
}

// We add between +1 and +2
func (c *circle) Add(m *marble) {
	// Next chain
	m.next = c.head.next.next
	c.head.next.next = m

	// Prev chain
	m.next.prev = m
	m.prev = c.head.next

	c.head = m
}

// Remove the marble -7
func (c *circle) Remove() int {
	curr := c.head
	for i := 0; i < 7; i++ {
		curr = curr.prev
	}

	curr.prev.next = curr.next
	curr.next.prev = curr.prev

	c.head = curr.next

	return curr.value
}

func play(noPlayers int, stop int) int64 {
	players := map[int]int64{}
	game := makeCircle()
	for m := 1; m <= stop; m++ {
		if m%23 == 0 {
			worth := game.Remove() + m
			player := (m - 1) % noPlayers
			players[player] += int64(worth)
		} else {
			game.Add(makeMarble(m))
		}
	}
	max := int64(0)
	for _, x := range players {
		if x > max {
			max = x
		}
	}
	return max
}

func main() {
	var noPlayers, stop int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if _, err := fmt.Sscanf(
			s.Text(),
			"%d players; last marble is worth %d points",
			&noPlayers,
			&stop,
		); err != nil {
			log.Fatal(err)
		}
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(play(noPlayers, stop * 100))
}
