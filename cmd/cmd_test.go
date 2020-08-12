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
			[]string{"add", "implement task deletion"},
			Param{
				"add",
				"implement task deletion",
			},
		},
		{
			[]string{"do", "1"},
			Param{
				"do",
				"1",
			},
		},
		{
			[]string{},
			Param{
				"",
				"",
			},
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
