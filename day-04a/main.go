package main

import (
	"time"
	"bufio"
	"os"
	"log"
	"fmt"
	"sort"
)

const (
	shiftStart = iota
	fallsAsleep = iota
	wakesUp = iota
)


type event struct {
	at time.Time
	eventType int
	guardID int
}

func makeEvent(line string) (event, error) {
	var eventTime time.Time
	var eventType int
	guardID := -1

	rawDate := line[1:17]
	eventDesc := line[19:]

	eventTime, err := time.Parse("2006-01-02 15:04", rawDate)
	if err != nil {
		return event{}, err
	}
	switch eventDesc[0] {
	case 'f':
		eventType = fallsAsleep
	case 'w':
		eventType = wakesUp
	case 'G':
		eventType = shiftStart
		if _, err := fmt.Sscanf(eventDesc, "Guard #%d begins shift", &guardID); err != nil {
			return event{}, err
		}
	default:
		return event{}, fmt.Errorf("Unrecognized event type: %s", eventDesc)
	}
	return event{eventTime, eventType, guardID}, nil
}

func main() {

	// Read events in order
	events := make([]event, 0)

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		evt, err := makeEvent(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		events = append(events, evt)
	}
	sort.Slice(events, func(i int, j int) bool {
		return events[i].at.Before(events[j].at)
	})

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	// Find the total number of minutes for each guard
	totalMinutes := make(map[int]int)
	minutesFreq := make(map[int][60]int)
	var guardID int
	var minStart int
	for _, evt := range events {
		switch(evt.eventType) {
		case shiftStart:
			guardID = evt.guardID
		case fallsAsleep:
			minStart = evt.at.Minute()
		case wakesUp:
			thisShift := evt.at.Minute() - minStart
			totalMinutes[guardID] += thisShift

			buf := minutesFreq[guardID]
			for i := minStart; i < evt.at.Minute(); i++ {
				buf[i]++
			}
			minutesFreq[guardID] = buf
		}
	}


	// Find max guard
	maxGuard, maxTotal := -1, -1
	for guardID, total := range totalMinutes {
		if total > maxTotal {
			maxTotal = total
			maxGuard = guardID
		}
	}


	// Find max minutes
	maxMinute, maxTimes := -1, -1
	for minute, times := range minutesFreq[maxGuard] {
		if times > maxTimes {
			maxTimes = times
			maxMinute = minute
		}
	}
	fmt.Println(maxMinute * maxGuard)
}
