package app

import "sort"

// 俄罗斯信封套娃问题
// 在二维平面里求最长递增子序列

//最长递增子序列问题
//首先按照宽度w从小到大排序，w相同的再按照h高从大到小排序
//之后以h作为一个数组，在这个数组上计算LIS的长度就是答案

//这个解法的关键在于，对于宽度 w 相同的数对，要对其高度 h 进行降序排序。
//因为两个宽度相同的信封不能相互包含的，逆序排序保证在 w 相同的数对中最多只选取一个

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		// 首先按照w升序排列， 当w相同时,按照h降序排列
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	f := []int{}
	//这块没看懂
	for _, e := range envelopes {
		h := e[1]
		if i := sort.SearchInts(f, h); i < len(f) {
			f[i] = h
		} else {
			f = append(f, h)
		}
	}
	return len(f)
}
