package main

import "fmt"

func deferRun() {
	a := 1
	b := 2
	defer deferCall("first_", a, deferCall("first", a, b))
	a = 0
	defer deferCall("second_", a, deferCall("second", a, b))

}

func deferCall(name string, x, y int) int {
	fmt.Println(name, x, y, x+y)
	return x + y
}

func deferReturn() int {
	i := 0
	addi := func() {
		i++
		fmt.Println("defer", i)
	}
	defer addi()
	defer addi()
	// 此处 return i 是值拷贝操作 将 i 理解为 return 函数的参数 // 等价于 s:=i; return s;
	return i
}

func main() {
	deferRun()
	return
}
