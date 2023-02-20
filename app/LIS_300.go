package app

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	for i := range nums {
		for j :=0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	}
	ans := 1
	for _, d := range dp {
		ans = max(ans, d)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func lengthOfLIS2(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := range nums {
		for j :=0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	ans := 1
	for i := range dp {
		ans = max(ans, dp[i])
	}
	return ans
}

//还有更加优化的版本，加入二分查找
func lengthOfLIS3(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	stack := []int{nums[0]}

	for i := 1; i<len(nums); i++ {
		if nums[i] > stack[len(stack)-1] {
			stack = append(stack, nums[i])
		} else if nums[i] < stack[len(stack)-1] {
			//开始做替换
			left, right := 0, len(stack)-1
			for left <=right {
				mid := left + (right-left)/2
				//找第一个大于nums[i]的值
				if nums[i] < stack[mid] {
					right = mid-1
				} else if nums[i] > stack[mid] {
					left = mid+1
				} else {
					left = mid
					break
				}
			}
			stack[left] = nums[i]
		}
	}
	return len(stack)
}