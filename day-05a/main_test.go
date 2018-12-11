package main

import (
	"testing"
	"reflect"
)

func TestSimplify(t *testing.T) {

	tables := []struct {
		name string
		subject *protein
		expected *protein
		simplified bool
	}{
		{
			"Empty protein",
			&protein{[]rune{}, 0},
			&protein{[]rune{}, 0},
			false,
		},
		{
			"a",
			&protein{[]rune{}, 'a'},
			&protein{[]rune{}, 'a'},
			false,
		},
		{
			"aa",
			&protein{[]rune{'a'}, 'a'},
			&protein{[]rune{'a'}, 'a'},
			false,
		},
		{
			"aA",
			&protein{[]rune{'a'}, 'A'},
			&protein{[]rune{}, 0},
			true,
		},
		{
			"Aa",
			&protein{[]rune{'A'}, 'a'},
			&protein{[]rune{}, 0},
			true,
		},
		{
			"baA",
			&protein{[]rune{'b', 'a'}, 'A'},
			&protein{[]rune{}, 'b'},
			true,
		},
		{
			"cbaA",
			&protein{[]rune{'c', 'b', 'a'}, 'A'},
			&protein{[]rune{'c'}, 'b'},
			true,
		},
		{
			"dcbaA",
			&protein{[]rune{'d', 'c', 'b', 'a'}, 'A'},
			&protein{[]rune{'d', 'c'}, 'b'},
			true,
		},
	}

	for _, table := range tables {
		res := table.subject.simplify()
		if res != table.simplified {
			t.Errorf("[%s] Actual: %v, expected: %v", table.name, res, table.simplified)
		}
		if !reflect.DeepEqual(table.subject, table.expected) {
			t.Errorf("[%s] Actual: %v, expected: %v", table.name, table.subject, table.expected)
		}
	}
}
