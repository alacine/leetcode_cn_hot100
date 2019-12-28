package q001

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	s := []int{2, 7, 11, 15}
	target := 9
	exp := []int{0, 1}
	got := twoSum(s, target)
	for i, num := range exp {
		fmt.Printf("exp[%d]=[%d], got[%d]=[%d]\n", i, exp[i], i, num)
		if num != got[i] {
			t.Errorf("exp[%d]=[%d], got[%d]=[%d]\n", i, exp[i], i, num)
		}
	}
}
