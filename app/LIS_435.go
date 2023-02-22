package app

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	//其实就是最长递增子序列的变种
	//只不过从数组变成了区间段
	if len(intervals) <= 1 {
		return 0
	}
	dp := make([]int, len(intervals))
	for i := range dp {
		dp[i] = 1
	}
	//先排序 快排
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := range intervals {
		for j :=0; j<i; j++ {
			a := intervals[i]
			b := intervals[j]
			if b[1] <= a[0] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	maxLen := 1
	for i := range dp {
		maxLen = max(maxLen, dp[i])
	}
	return len(intervals) - maxLen
}

//会超时，那么采用二分查找加单调栈才能过吧
func eraseOverlapIntervals2(intervals [][]int) int {
	//其实就是最长递增子序列的变种
	//只不过从数组变成了区间段
	if len(intervals) <= 1 {
		return 0
	}

	//先排序 快排
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	return 0

}

//方法三，贪心算法
//核心是按照区间右侧进行排序
//右侧越小的一定越先结束
//非常trick也非常常见
func eraseOverlapIntervals3(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	//最多的不重叠区间个数
	cnt := 0
	end := intervals[0][1]
	for i:=1; i<len(intervals);i++ {
		if end <= intervals[i][0] {
			end = intervals[i][1]
		} else {
			cnt++
		}
	}
	return cnt
}