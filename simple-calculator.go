package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func div(a, b int) int {
	return a / b
}
func main() {
	calculator := make(map[string]func(int, int) int)
	calculator["add"] = add
	calculator["sub"] = sub
	calculator["mul"] = mul
	calculator["div"] = div
	op := os.Args[1]
	num1Str := os.Args[2]
	num2Str := os.Args[3]
	if len(os.Args) != 4 {
		fmt.Println("格式错误 \n 正确格式: go run demo1.go [操作][操作数1] [操作数2]")
		return
	}
	num1, err1 := strconv.Atoi(num1Str)
	num2, err2 := strconv.Atoi(num2Str)
	if err1 != nil || err2 != nil {
		fmt.Println("错误，操作数为整数")
		return
	}
	fn, ok := calculator[op]
	if ok {
		fmt.Println(fn(num1, num2))
	} else {
		fmt.Println("错误，操作必须为(add sub mul div)")
		return
	}
}
+
