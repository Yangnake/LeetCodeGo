package app

//给定⼀个按照升序排列的整数数组 nums，和⼀个⽬标值 target，找出给定⽬标值在数组中的开始位置和结
//束位置。如果数组中不存在⽬标值 target，返回 [-1, -1]

/*
示例1:
输⼊：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

示例2:
输⼊：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
 */
// 首先，二分查找有闭区间，开闭区间，闭开区间是那种方式去书写
// 闭区间就是 left <= right ; 开闭区间为left < right
// 循环不变量，left的左侧一定是小于target的值，right的右侧一定是大于target的值
// 通过循环不变量即可判断出边界位置

func lower_bound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right { //[left, right] 闭区间
		mid := left + (right-left)/2
		if nums[mid] < target {
			//我们需要询问的区间缩小到[mid+1, right]
			left = mid + 1
		} else {
			//需要询问的区间变为[left, mid-1]
			right = mid - 1
		}
	}
	return left
}

func lower_bound2(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right { //左闭右开区间[left, right)
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1 // [mid+1, right)
		} else {
			right = mid // [left, mid)
		}
	}
	return left // or right
}

/*
>= --> >= 大于等于x
> --> >=  将大于转化为大于等于x+1 (数组都是整数时可用)
< --> 可以看成大于等于x -1
<= --> 可以看成大于x -1
 */
func searchRange(nums []int, target int) []int {
	left := lower_bound(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1,-1}
	}
	//如果left存在那么right一定存在
	right := lower_bound(nums, target+1)-1
	return []int{left, right}
}

// 一个简单的题目：搜索排序数组中的target,如果存在返回其索引，如果不存在返回它应该被插入的位置
// 其实就是二分查找 深刻理解循环不变量的理论，任何二分问题都能轻松解决
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right { // [left, right]
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1 //依据循环不变量的理论，left的左侧始终小于target
		} else {
			right = mid - 1 //right的右侧始终大于等于target
		}
	}
	return left
}