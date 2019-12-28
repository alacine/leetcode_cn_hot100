package q003

import (
	"bytes"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	a, b := []byte{}, []byte{}
	t := []byte(s)
	for _, char := range t {
		la, lb := len(a), len(b)
		if ib := bytes.IndexByte(b, char); ib != -1 {
			if la < lb {
				a = b
			}
			b = b[ib+1:]
			b = append(b, char)
		} else {
			b = append(b, char)
		}
	}
	return max(len(a), len(b))
}
