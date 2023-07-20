package productexceptself

func productExceptSelf(nums []int) []int {
	n := len(nums)

	var pre []int
	var nxt []int

	pre[0] = 1
	nxt[n-1] = 1

	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] * nums[i-1]
	}

	for n := n - 2; n >= 0; n-- {
		nxt[n] = nxt[n+1] * nums[n+1]
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = pre[i] * nxt[i]
	}
	return nums
}
