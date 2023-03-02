package backtrack

func permute(nums []int) [][]int {
	//典型回溯法
	//初始化
	n := len(nums)
	//初始化结果集
	res := make([][]int, 0)
	//初始化标志位数组
	used := make([]bool, n)
	//初始化单个结果集
	path := make([]int, 0)

	//递归函数，从树的根结点开始，逐层向下去叶子节点中寻找单个结果集，并加入到最终结果集中
	var dfs func(int, []int, []bool)
	//depth：树的层数，逐层向下找
	//path单个结果集，记录树中一条路径中从根结点到叶子节点的变化过程
	//used：标志位数组，用来标记到树中的某个分支的某一层时，已经走过的边
	dfs = func(depth int, path []int, used []bool) {
		if depth == n {
			newPath := make([]int, 0)
			newPath = append(newPath, path...)
			res = append(res, newPath)
			return
		}
		for i:=0; i<n;i++ {
			if !used[i] {
				path = append(path, nums[i])
				used[i] = true
				dfs(depth+1, path, used)
				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}
	dfs(0, path, used)
	return res
}

func permute2(nums []int) [][]int {
	//全排列问题，就是每次从队列中选择一个元素，当把队列中的全部元素选出则生成一个排列
	res := [][]int{}
	used := make([]bool, len(nums))
	path := []int{}
	//路径， 选择列表，记录访问过的元素列表
	var dfs func([]int)

	dfs = func(nums []int) {
		if len(path) == len(nums) {
			tmp := append([]int{}, path...)
			res = append(res, tmp)
			return
		}
		for i := 0; i<len(nums); i++ {
			if !used[i] {
				path = append(path, nums[i])
				used[i] = true
				dfs(nums)
				path = path[:len(path)-1]
				used[i] = false
			}
		}
		return
	}
	dfs(nums)
	return res
}