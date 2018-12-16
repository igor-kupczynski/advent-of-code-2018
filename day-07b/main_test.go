package main

import (
	"testing"
)

func TestTakes(t *testing.T) {
	tables := []struct {
		stepName string
		takesTime int
	}{
		{"A", 61},
		{"Z", 86},
	}
	for _, tab := range tables {
		took := takes(tab.stepName)
		if took != tab.takesTime {
			t.Errorf("%s should take %d, but took %d\n", tab.stepName, tab.takesTime, took)
		}
	}
}
