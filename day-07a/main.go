package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type step struct {
	name     string
	previous []string
	next     []string
}

type state struct {
	done          []string
	remaining     []string
	prerequisites map[string][]string
	nexts         map[string][]string
}

func buildState(spec map[string]step) state {
	done := []string{}
	remaining := []string{}
	prereqs := map[string][]string{}
	nexts := map[string][]string{}

	for k, v := range spec {
		remaining = append(remaining, k)

		prerequisites := make([]string, len(v.previous))
		copy(prerequisites, v.previous)
		prereqs[k] = prerequisites

		next := make([]string, len(v.next))
		copy(next, v.next)
		nexts[k] = next
	}
	sort.Sort(sort.StringSlice(remaining))

	return state{done, remaining, prereqs, nexts}
}

func (s *state) IsDone() bool {
	return len(s.remaining) == 0
}

func (s *state) isAvailable(step string) bool {
	return len(s.prerequisites[step]) == 0
}

func (s *state) Next() {
	for idx, step := range s.remaining {
		if s.isAvailable(step) {
			s.done = append(s.done, step)
			s.remaining = append(s.remaining[:idx], s.remaining[idx+1:]...)
			for _, next := range s.nexts[step] {
				p := s.prerequisites[next]
				for idx, item := range p {
					if item == step {
						p = append(p[:idx], p[idx+1:]...)
						break
					}
				}
				s.prerequisites[next] = p
			}
			break
		}
	}
}

func (s *state) Report() string {
	return strings.Join(s.done, "")
}

func main() {
	spec := map[string]step{}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var from, to string
		fmt.Sscanf(s.Text(), "Step %s must be finished before step %s can begin.", &from, &to)

		if _, exists := spec[from]; !exists {
			spec[from] = step{from, []string{}, []string{}}
		}
		fromStep := spec[from]
		fromStep.next = append(fromStep.next, to)
		spec[from] = fromStep

		if _, exists := spec[to]; !exists {
			spec[to] = step{to, []string{}, []string{}}
		}
		toStep := spec[to]
		toStep.previous = append(toStep.previous, from)
		spec[to] = toStep
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	state := buildState(spec)

	for !state.IsDone() {
		state.Next()
	}
	fmt.Println(state.Report())
}
