package main

import "fmt"

func main() {
	// 配列B を受け取って要素の 2倍にした配列を返す関数Double(B) を作成する
	B := []int{1, 9, 3, 5}
	fmt.Printf("Double(B) = %v\n", Double(B))
	// 出力例: "Double(B) = [2 18 6 10]"
}
