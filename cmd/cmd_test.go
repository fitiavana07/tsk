package cmd

import (
	"testing"
)

func TestParseArgs(t *testing.T) {
	type Case struct {
		args     []string
		expected Param
	}

	cases := []Case{
		{
			// right args
			[]string{"add", "implement task deletion"},
			Param{"add", "implement task deletion"},
		},
		{
			// action and number as arg
			[]string{"do", "1"},
			Param{"do", "1"},
		},
		{
			// no action, no arg
			[]string{},
			Param{"", ""},
		},
		{
			// too many args
			[]string{"skip", "too", "many", "args"},
			Param{"", ""},
		},
	}

	for _, c := range cases {
		got := ParseArgs(c.args)
		if got != c.expected {
			t.Errorf("FAIL: parseFlags(%v) == %v, expected %v", c.args, got, c.expected)
		} else {
			t.Logf("SUCCESS: parseFlags(%v) == %v, expected %v", c.args, got, c.expected)
		}
	}
}
