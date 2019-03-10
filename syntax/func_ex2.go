package main

import "fmt"

func main() {
	// 配列Bを受け取って要素の2倍にした配列を返す関数Double(B)を作成する
	B := []int{1, 9, 3, 5}
	fmt.Printf("Double(B) = %v\n", Double(B))
	// 出力: "Double(B) = [2 18 6 10]"
}

func Double(B []int) []int {
	C := make([]int, len(B))
	for i, e := range B {
		C[i] = 2 * e
	}
	return C
}
