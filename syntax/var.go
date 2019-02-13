package main

import "fmt"

func main() {
	// 変数の宣言
	// var 変数名 型名 の形で宣言する
	var v1 int

	// 宣言された直後は0値と呼ばれる値に初期化される
	fmt.Printf("v1 = %d\n\n", v1)

	// 変数への値の代入
	// 変数名 = 値
	v1 = 1
	fmt.Printf("v1 = %d\n", v1)
	v1 = 2
	fmt.Printf("v1 = %d\n\n", v1)

	// 値の基本型と0値
	var b1 bool
	var i1 int
	var s1 string
	var f1 float32
	fmt.Printf("b1 = %t, i1 = %d, s1 = '%s', f1 = %f\n\n", b1, i1, s1, f1)

	// 変数の宣言と代入を同時に行う
	// "" で囲まれた値はstring(文字列)として解釈される
	v2 := "v2"
	fmt.Printf("v2 = '%s'\n\n", v2)

	// 既存の変数を再宣言しようとするとエラー
	// v1 := 3

	// 宣言した型と違う値を代入してもエラー
	// v1 = "v1"

	// 定数はconstで宣言できる
	const c1 = 1
	fmt.Printf("c1 = %d\n\n", c1)

	// 定数値は値を変更できない
	// c1 = 2

	// 変数はスコープごとで別に記録される
	{
		// ここではv1は外側のスコープの変数の値を参照する
		fmt.Printf("v1 = %d\n", v1)
		// スコープ内の変数を定義
		v1 := 100
		fmt.Printf("v1 = %d\n", v1)
	}
	// スコープ内での新変数への操作は影響しない
	fmt.Printf("v1 = %d\n\n", v1)

	// 応用構文
	// 複数の同じ型の変数を同時に宣言
	var v3, v4 int

	// 同時複数の変数に代入
	v3, v4 = 3, 4
	fmt.Printf("v3 = %d, v4 = %d\n\n", v3, v4)

	// 同時に複数の型の変数を宣言
	var (
		vg1 int
		vg2 string
		vg3 float32
	)
	vg1, vg2, vg3 = 1, "2", 3.0
	fmt.Printf("vg1 = %d, vg2 = %s, vg3 = %f\n", vg1, vg2, vg3)
}
