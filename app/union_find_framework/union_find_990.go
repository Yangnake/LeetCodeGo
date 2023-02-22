package union_find_framework

/*
给定一个由表示变量之间关系的字符串方程组成的数组，每个字符串方程 equations[i] 的长度为 4，并采用两种不同的形式之一："a==b" 或"a!=b"。在这里，a 和 b 是小写字母（不一定不同），表示单字母变量名。
只有当可以将整数分配给变量名，以便满足所有给定的方程时才返回true，否则返回 false。 

示例 1：

输入：["a==b","b!=a"]
输出：false
解释：如果我们指定，a = 1 且 b = 1，那么可以满足第一个方程，但无法满足第二个方程。没有办法分配变量同时满足这两个方程。
示例 2：

输入：["b==a","a==b"]
输出：true
解释：我们可以指定 a = 1 且 b = 1 以满足满足这两个方程。
示例 3：

输入：["a==b","b==c","a==c"]
输出：true
示例 4：

输入：["a==b","b!=c","c==a"]
输出：false
示例 5：

输入：["c==c","b==d","x!=z"]
输出：true
 */

func equationsPossible(equations []string) bool {
	//这是一个考察并查集的题目
	parent := make([]int, 26)
	for i := 0; i<26; i++ {
		parent[i] = i
	}
	for _, str := range equations {
		if str[1] == '=' {
			//将等于的两字数字放入一棵树中
			index1 := int(str[0]-'a')
			index2 := int(str[3]-'a')
			union(parent, index1, index2)
		}
	}
	for _, str := range equations {
		if str[1] == '!' {
			index1 := int(str[0] - 'a')
			index2 := int(str[3] - 'a')
			//对不相等的两个数查找并查集，看是否联通，如果联通则不满足条件
			if find(parent, index1) == find(parent, index2) {
				return false
			}
		}
	}
	return true
}

//将两棵树合并
func union(parent []int, index1, index2 int) {
	parent[find(parent, index1)] = find(parent, index2)
}

//一直向上找到根节点
func find(parent []int, index int) int {
	for parent[index] != index {
		parent[index] = parent[parent[index]]
		index = parent[index]
	}
	return index
}