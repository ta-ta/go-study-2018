package main

import "fmt"

func main() {
	// 課題1: 0 から 99 までの数字を順に追加したスライスnumsを作成する
	nums := make([]int, 0, 100)

	fmt.Printf("nums = %v\n", nums)

	// 課題2: (課題1から) 最後の要素から逆順にスライスの内容を出力
	// 出力: 99 98 .... 3 2 1 0

	// 課題3: (課題1から) 奇数番目の要素のみを出力
	// 出力: 1 3 5 ... 97 99
}
