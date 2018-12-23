package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type five struct {
	pattern [5]bool
	out     bool
}

func (f *five) matches(plants []bool) bool {
	for i := 0; i < 5; i++ {
		if plants[i] != f.pattern[i] {
			return false
		}
	}
	return true
}

func newFive(line string) *five {
	var pattern [5]bool
	for i := 0; i < 5; i++ {
		pattern[i] = hasPlant(line[i])
	}

	out := hasPlant(line[9])
	return &five{pattern, out}
}

func hasPlant(c byte) bool {
	var hasPlant bool
	switch c {
	case '.':
		hasPlant = false
	case '#':
		hasPlant = true
	default:
		log.Fatalln("Can't parse plant description", c)
	}
	return hasPlant
}

type fives []*five

func (f *fives) next(plants []bool) bool {
	for _, five := range []*five(*f) {
		if five.matches(plants) {
			return five.out
		}
	}
	return false
}

func (f *fives) nextGen(curr generation) generation {
	next := generation{make([]bool, len(curr.plants)), curr.center}
	for i := 2; i < len(curr.plants)-2; i++ {
		next.plants[i] = f.next(curr.plants[i-2 : i+3])
	}
	return next.shrink()
}

type generation struct {
	plants []bool
	center int
}

func (g generation) String() string {
	var b strings.Builder
	for i, x := range g.plants {
		if i == g.center {
			b.WriteByte('|')
		}
		if x {
			b.WriteByte('#')
		} else {
			b.WriteByte('.')
		}
	}
	return b.String()
}

func (g generation) shrink() generation {
	var start, stop = 0, len(g.plants)
	for i, x := range g.plants {
		if x || i == g.center {
			start = i
			break
		}
	}

	for i := len(g.plants) - 1; i >= 0; i-- {
		x := g.plants[i]
		if x || i == g.center {
			stop = i + 1
			break
		}
	}

	// extra two empty pots in front and in the back
	newPlants := append([]bool{false, false, false, false}, g.plants[start:stop]...)
	newPlants = append(newPlants, false, false, false, false)

	var shrinked generation
	shrinked.center = g.center - start + 4 // 4 because of the two extra empty pots
	shrinked.plants = newPlants
	return shrinked
}

func (g generation) sum() int {
	sum := 0
	for i, x := range g.plants {
		if x {
			sum += i - g.center
		}
	}
	return sum
}

func readInput() (f fives, gen generation) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "initial state: ") {
			for i := 15; i < len(line); i++ {
				gen.plants = append(gen.plants, hasPlant(line[i]))
			}
			gen = gen.shrink()
		} else {
			f = append(f, newFive(line))
		}
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func main() {
	sampleEnd := 1000
	samples := make([]int, sampleEnd+1)
	fives, generation := readInput()
	samples[0] = generation.sum()
	for i := 1; i <= sampleEnd; i++ {
		generation = fives.nextGen(generation)
		samples[i] = generation.sum()

	}

	// Hyphotesis
	ax := samples[sampleEnd] - samples[sampleEnd-1]
	b := samples[sampleEnd] - ax * sampleEnd

	// Verification
	for i := sampleEnd - 10; i <= sampleEnd; i++ {
		if expected := ax * i + b; expected != samples[i] {
			log.Fatalf("Failed verification, expected=%d, actual=%d\n", expected, samples[i])
		}
	}

	fmt.Printf("sum[x] = %d*x +%d\n", ax, b)


	// Task
	generations := int64(50000000000)
	fmt.Println(int64(ax) * generations + int64(b))
}
