package bfs_framework

func openLock(deadends []string, target string) (step int) {
	deads := map[string]bool{}
	for _, d := range deadends {deads[d] = true}
	if deads["0000"] == true {return -1} // 非法
	visited := map[string]bool{"0000": true}
	//通过切片可以很方便地模拟队列的操作
	q := []string{"0000"}
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ { //将每一层遍历完
			v := q[0]
			q = q[1:]
			if v == target {return}
			for j := 0; j < 4; j++ { // 四位数字
				up, down := plusOne(v, j), minusOne(v, j)
				if !visited[up] && !deads[up] {q = append(q, up); visited[up] = true} // 避免重复循环访问已访问过的元素，若属于deads则应剪枝
				if !visited[down] && !deads[down] {q = append(q, down); visited[down] = true}
			}
		}
		step++
	}
	return -1
}

func plusOne(s string, idx int) string {
	bs := []byte(s) //先转为byte再转为string
	if bs[idx] == '9' {bs[idx] = '0'} else {bs[idx]++}
	return string(bs)
}

func minusOne(s string, idx int) string {
	bs := []byte(s)
	if bs[idx] == '0' {bs[idx] = '9'} else {bs[idx]--}
	return string(bs)
}

package bfs_framework

func openLock2(deadends []string, target string) int {
	if target == "0000" { //非法case
		return -1
	}
	deads := map[string]bool{}
	for _, d := range deadends {
		deads[d] = true
	}
	visited := map[string]bool{"0000": true}
	que := []string{"0000"}
	step := 0
	for len(que) > 0 {
		sz := len(que)
		for i := 0; i < sz; i++ {
			v := que[0]
			que = que[1:]
			if v == target {
				return step
			}
			for j := 0; j < 4; j++ {
				up, down := plusOne(v, j), minusOne(v, j)
				if !visited[up] && !deads[up] {
					que = append(que, up)
					visited[up] = true
				}
				if !visited[down] && !deads[down] {
					que = append(que, up)
					visited[down] = true
				}
			}
		}
	}
	//走到这里说明能走的路全部走完了，说明没有解
	return -1
}
