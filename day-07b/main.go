package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const (
	baseTaskTime = 60
	totalWorkers = 5
)

type step struct {
	name     string
	previous []string
	next     []string
}

func takes(step string) int {
	return baseTaskTime + 1 + int(step[0]-'A')
}

type state struct {
	tick          int
	done          []string
	remaining     []string
	prerequisites map[string][]string
	nexts         map[string][]string
	workerTask    []string
	workerDone    []int
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

	return state{
		tick:          -1,
		done:          done,
		remaining:     remaining,
		prerequisites: prereqs,
		nexts:         nexts,
		workerTask:    make([]string, totalWorkers),
		workerDone:    make([]int, totalWorkers),
	}
}

func (s *state) IsDone() bool {
	for _, wt := range s.workerTask {
		if wt != "" {
			return false
		}
	}
	return len(s.remaining) == 0
}

func (s *state) isAvailable(step string) bool {
	return len(s.prerequisites[step]) == 0
}

func (s *state) Next() string {
	for idx, step := range s.remaining {
		if s.isAvailable(step) {
			s.remaining = append(s.remaining[:idx], s.remaining[idx+1:]...)
			return step
		}
	}
	return ""
}

func (s *state) markDone(workerIdx int) {
	step := s.workerTask[workerIdx]
	s.done = append(s.done, step)
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
	s.workerTask[workerIdx] = ""
}

func (s *state) MaybeMarkDone() {
	for i := 0; i < totalWorkers; i++ {
		if s.workerTask[i] != "" && s.workerDone[i] == s.tick {
			s.markDone(i)
		}
	}
}

func (s *state) IsWorkerFree(workerIdx int) bool {
	return s.workerTask[workerIdx] == ""
}

func (s *state) Assign(workerIdx int, task string) {
	s.workerTask[workerIdx] = task
	s.workerDone[workerIdx] = s.tick+takes(task)
}

func (s *state) Tick() {
	s.tick++
}

func (s *state) Report() (int, string) {
	return s.tick, strings.Join(s.done, "")
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
		// Tick
		state.Tick()

		// Check if anything should be marked as done
		state.MaybeMarkDone()

		// Check if we have any free workers
		for i := 0; i < totalWorkers; i++ {
			if state.IsWorkerFree(i) {
				// If so, check if we have a task for them to assign
				if next := state.Next(); next != "" {
					state.Assign(i, next)
				}
			}
		}
	}
	fmt.Println(state.Report())
}
