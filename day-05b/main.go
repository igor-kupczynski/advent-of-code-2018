package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	reactionDist = 'A' - 'a'
)

type protein struct {
	prev       []rune
	curr       rune
	charFilter func(ch rune) bool
}

func emptyProtein(typeToSkip rune) protein {
	return protein{
		[]rune{},
		0,
		func(ch rune) bool { return ch != typeToSkip && ch != (typeToSkip+reactionDist) },
	}
}

func (p *protein) feed(next rune) {
	if p.charFilter(next) {
		if p.curr != 0 {
			p.prev = append(p.prev, p.curr)
		}
		p.curr = next
		for p.simplify() {
		}
	}
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

func allTypes(line string) map[rune]int {
	result := map[rune]int{}
	for _, x := range strings.ToLower(line) {
		result[x] = 0
	}
	return result

}

func main() {

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		types := allTypes(line)
		for t := range types {
			p := emptyProtein(t)
			for _, x := range line {
				p.feed(x)
			}
			types[t] = len(p.String())
		}

		minLength := 10762  // From the day-05a :)
		for _, l := range types {
			if l < minLength {
				minLength = l
			}
		}
		fmt.Println(minLength)
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
