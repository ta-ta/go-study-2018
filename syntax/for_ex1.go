package main

import "fmt"

func main() {
	// 0 から 99 まで計100個の数字を順に格納したスライスnums を作成する
	nums := make([]int, 100)         // 修正箇所
	for i := 0; i < len(nums); i++ { // 修正箇所
		nums[i] = i // 修正箇所
	}
	fmt.Printf("nums = %v\n", nums)
	fmt.Println()
	fmt.Println()

	// nums を最後の要素から順に出力
	// 出力: 99 98 .... 3 2 1 0
	for j := len(nums) - 1; j >= 0; j-- { // 修正箇所
		fmt.Println(nums[j])
	}

	/*
		// 奇数番目の要素のみを出力
		// 出力: 1 3 5 ... 97 99
		for k := 1; k < 100; k += 2 {
			fmt.Println(nums[k])
		}
	*/
}
