package main

import (
	"bufio"
	"os"
	"log"
	"fmt"
)

type protein struct {
	prev []rune
	curr rune
}

func emptyProtein() protein {
	return protein{[]rune{}, 0}
}

const (
	reactionDist = 'a' - 'A'
)

func (p *protein) feed(next rune) {
	if p.curr != 0 {
		p.prev = append(p.prev, p.curr)
	}
	p.curr = next
	for p.simplify() {}
}

func (p *protein) simplify() bool {
	if len(p.prev) == 0 {
		return false
	}
	last := len(p.prev) - 1
	dist := p.curr - p.prev[last]
	if dist == reactionDist || dist == -reactionDist {
		if len(p.prev) < 2 {
			p.curr = 0
			p.prev = p.prev[:0]
		} else {
			p.curr = p.prev[last-1]
			p.prev = p.prev[:last-1]
		}
		return true
	}
	return false
}

func (p *protein) String() string {
	if p.curr == 0 {
		return ""
	}
	return string(append(p.prev, p.curr))
}

func main() {

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		p := emptyProtein()
		for _, x := range line {
			p.feed(x)
		}
		fmt.Println(len(p.String()))
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}


}
