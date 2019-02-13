package main

import "fmt"

func main() {
	// 課題1: 0 から 99 までの数字を順に追加したスライスnumsを作成する
	nums := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("nums = %v\n", nums)

	// 課題2: (課題1から) 最後の要素から逆順にスライスの内容を出力
	// 出力: [99 98 .... 3 2 1 0]
	for j := len(nums) - 1; j >= 0; j-- {
		fmt.Println(nums[j])
	}

	// 課題3: (課題1から) 奇数番目の要素のみを出力
	// 出力: [1 3 5 ... 97 99]
	// ** 回答1
	k := 1
	for k < 100 {
		fmt.Println(nums[k])
		k += 2
	}
	// ** 回答2
	for k := 1; k < 100; k += 2 {
		fmt.Println(nums[k])
	}
}
