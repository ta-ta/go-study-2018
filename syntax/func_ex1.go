package main

import "fmt"

func main() {
	// 0からNまでの和を計算する関数Sum(N)を作成する
	fmt.Printf("Sum(100) = %d\n", Sum(100))
	// 出力: "Sum(100) = 5050"
}

// 0 から n までの数字の和を計算する関数
func Sum(n int) int {
	result := 0
	for i := 0; i <= n; i++ {
		result += i
	}
	return result
}
