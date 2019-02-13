package main

import "fmt"

func main() {
	// 課題: 配列Bを受け取って要素の2倍にした配列を返す関数Double(B)を実装
	B := []int{1, 9, 3, 5}
	fmt.Printf("Double(B) = %v\n", Double(B))
	// 出力: "Double(B) = [2 18 6 10]"
}
