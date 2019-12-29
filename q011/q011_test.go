package q011

import "testing"

func Test(t *testing.T) {
	var tests = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	var exp int = 49
	got := maxArea(tests)
	if exp != got {
		t.Errorf("exp = %d, got = %d\n", exp, got)
	}
}
