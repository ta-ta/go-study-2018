package main

import "fmt"

func main() {
	// slice の宣言 (`[]int`型)
	var s1 []int

	// 初期値
	fmt.Println("初期値")
	fmt.Printf("s1 = %v\n", s1)
	fmt.Println()
	fmt.Println()

	// 代入
	s1 = []int{1, 2, 3}
	fmt.Println("代入")
	fmt.Printf("s1 = %v\n", s1)
	fmt.Println()
	fmt.Println()

	// スライスの要素数を調べる
	fmt.Println("要素数")
	fmt.Printf("len(s1) = %d\n", len(s1))
	fmt.Println()
	fmt.Println()

	// スライスの特定の要素を参照する
	fmt.Println("参照")
	fmt.Printf("s1[1] = %d\n", s1[1])
	// 要素数以上のインデックスを参照するとエラー
	// fmt.Printf("s1[100] = %d\n", s1[100])
	fmt.Println()
	fmt.Println()

	// スライスに要素を追加する
	fmt.Println("追加")
	s2 := append(s1, 4) // s1 の末尾に 4 を追加
	fmt.Printf("s2 = %v\n", s2)
	fmt.Println()
	fmt.Println()

	// スライスを連結する
	fmt.Println("連結")
	s12 := append(s1, s2...) // s1 と s2 を連結
	fmt.Printf("s12 = %v\n", s12)
	fmt.Println()
	fmt.Println()

	// 要素数を予約してスライスを作成する(make)
	// - 第一引数は, スライスの型
	// - 第二引数は, 要素数
	s3 := make([]int, 10)
	fmt.Println("make")
	fmt.Printf("len(s3) = %d, s3 = %v\n", len(s3), s3)
	fmt.Println()
	fmt.Println()

	/*
		// - 第三引数は 容量 (省略可)
		// 容量はあらかじめ確保しておくメモリ量, 大きいスライスを作成することがわかっている場合は
		// ここの数字をそれに近い数字にすると動作がはやくなる
		s4 := make([]int, 0, 10)
		fmt.Printf("len(s4) = %d, cap(s4) = %d, s4 = %v\n", len(s4), cap(s4), s4)
		fmt.Println()
		fmt.Println()

		// スライスから部分スライスをとりだす
		s5 := s2[1:3]
		fmt.Printf("s5 = %v\n", s5)
		fmt.Println()
		fmt.Println()

		// 取り出した部分スライスは、その部分の値を参照しているだけなので
		// 書き換えると元の配列も変更される
		s5[0] = 9
		fmt.Printf("s5 = %v\n", s5)
		fmt.Printf("s2 = %v\n", s2)
		fmt.Println()
		fmt.Println()

		// 元の部分スライスのコピーを作りたい場合はcopy
		s6 := make([]int, 2)
		copy(s6, s2[1:3]) // s3[1:3] から s6 へコピー

		s6[0] = 10
		fmt.Printf("s6 = %v\n", s6)
		fmt.Printf("s2 = %v\n", s2)
		fmt.Println()
		fmt.Println()
	*/
}
