package q001

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	ans := make([]int, 0)
	for i, num := range nums {
		_, isChecked := dict[target-num]
		if isChecked {
			ans = append(ans, dict[target-num])
			ans = append(ans, i)
			break
		}
		dict[num] = i
	}
	return ans
}
