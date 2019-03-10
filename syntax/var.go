package main

import "fmt"

func main() {
	// 基本型と変数宣言
	// var 変数名 型名
	var i1 int
	var f1 float32
	var s1 string
	var b1 bool

	// 宣言された直後は0値に初期化される
	fmt.Printf("宣言と初期化\n")
	fmt.Printf("int型の変数 i1 = %d\n", i1)
	fmt.Printf("float型の変数 f1 = %f\n", f1)
	fmt.Printf("string型の変数 s1 = %s\n", s1)
	fmt.Printf("bool型の変数 b1 = %t\n", b1)
	fmt.Printf("\n\n")

	// 変数へ値の代入
	// 変数名 = 値
	i1 = 10
	fmt.Printf("代入\n")
	fmt.Printf("i1 = %d\n", i1)
	fmt.Printf("\n\n")

	// 変数の宣言と代入を同時に行う
	// "" で囲まれた値はstring(文字列)として解釈される
	s2 := "s2"
	fmt.Printf("宣言と代入\n")
	fmt.Printf("s2 = '%s'\n", s2)
	fmt.Printf("\n\n")

	// 応用構文
	// 複数の同じ型の変数を同時に宣言
	var i2, i3 int
	fmt.Printf("i2 = %d, i3 = %d\n", i2, i3)

	// 複数の型の変数を宣言
	var (
		vg1 int
		vg2 string
		vg3 float32
	)
	vg1, vg2, vg3 = 1, "2", 3.0
	fmt.Printf("vg1 = %d, vg2 = %s, vg3 = %f\n", vg1, vg2, vg3)

	/*
		// 既存の変数を再宣言しようとするとエラー
		// i1 := 200

		// 宣言した型と違う値を代入してもエラー
		// i1 = "300"

		// 定数はconstで宣言できる
		const const_i1 = 1
		fmt.Printf("定数\n")
		fmt.Printf("const_i1 = %d\n", const_i1)
		fmt.Printf("\n\n")
		// 定数値は値を変更できない
		// c1 = 2

		// 変数はスコープごとに区別して記録される
		fmt.Printf("スコープ\n")
		{
			// ここではi1は外側のスコープの変数の値を参照する
			fmt.Printf("i1 = %d\n", i1)
			// スコープ内の変数を定義
			i1 := 100
			fmt.Printf("i1 = %d\n", i1)
		}
		// スコープ内での新変数への操作は影響しない
		fmt.Printf("i1 = %d\n", i1)
		fmt.Printf("\n\n")
	*/
}
