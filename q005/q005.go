package q005

func longestPalindrome(s string) string {
	l, maxlen, ans := len(s), 0, ""
	for i := 0; i < l; i++ {
		for j := 0; i-j >= 0 && i+j < l; j++ {
			if s[i-j] == s[i+j] {
				if j*2+1 > maxlen {
					maxlen = j*2 + 1
					ans = s[i-j : i+j+1]
				}
			} else {
				break
			}
		}
		for j, k := i, i+1; j >= 0 && k < l; j, k = j-1, k+1 {
			if s[j] == s[k] {
				if k-j+1 > maxlen {
					maxlen = k - j
					ans = s[j : k+1]
				}
			} else {
				break
			}
		}
	}
	return ans
}
