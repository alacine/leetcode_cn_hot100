package q003

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		s   string
		exp int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}
	for _, c := range tests {
		got := lengthOfLongestSubstring(c.s)
		if got != c.exp {
			t.Errorf("string = %q, exp = %d, got = %d\n", c.s, c.exp, got)
		}
	}
}
