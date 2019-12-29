package q005

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		s   string
		exp string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
	}
	for _, c := range tests {
		got := longestPalindrome(c.s)
		if c.exp != got {
			t.Errorf("exp = %q, got = %q\n", c.exp, got)
		}
	}
}
