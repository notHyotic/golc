package dp

// leetcode #213
func rob(nums []int) int {
    n := len(nums)
	if n == 1 {
		return nums[0]
	}

	n1 := nums[0 : n-1]
	n2 := nums[1 : n]

	return max(robHelper(n1), robHelper(n2))
}

func robHelper(nums []int) int {
	r1, r2 := 0, 0

	for _, n := range nums {
		tmp := max(r1+n, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
