package q011

func maxArea(height []int) int {
	ans := 0
	for i, j := 0, len(height)-1; i != j; {
		if height[i] < height[j] {
			if tmp := height[i] * (j - i); ans < tmp {
				ans = tmp
			}
			i++
		} else {
			if tmp := height[j] * (j - i); ans < tmp {
				ans = tmp
			}
			j--
		}
	}
	return ans
}
