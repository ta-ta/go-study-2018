package main

import (
	"fmt"
)

func main() {
	// 引数なし, 返り値なし関数 hello()
	fmt.Println("引数なし, 返り値なし関数")
	hello()
	fmt.Println()
	fmt.Println()

	// 引数あり, 返り値なし関数 hello2()
	fmt.Println("引数あり, 返り値なし関数")
	hello2("Hello World2")
	hello2("Hello World!")
	fmt.Println()
	fmt.Println()

	// 引数あり, 返り値あり関数 add()
	fmt.Println("引数あり, 返り値あり関数")
	v1 := add(1, 2)
	fmt.Printf("add(1, 2) = %d\n", v1)
	fmt.Println()
	fmt.Println()

	// 引数あり, 返り値が複数の関数 head_tail()
	fmt.Println("引数あり, 返り値が複数の関数")
	v2, v3 := head_tail([]int{3, 1, 4, 1, 5})
	fmt.Printf("head_tail([]int{3, 1, 4, 1, 5}) = %d, %d\n", v2, v3)
	fmt.Println()
	fmt.Println()

	/*
		// 関数div
		ret, err := div(10, 0)
		if err != nil {
			fmt.Printf("err = '%v'\n", err)
		}
		fmt.Printf("ret = %d\n\n", ret)

		ret, err = div(10, 2)
		if err != nil {
			fmt.Printf("err = '%v'\n", err)
		}
		fmt.Printf("ret = %d\n\n", ret)

		// ローカル関数の定義
		f1 := func(a, b int) int {
			return a + b
		}
		fmt.Printf("f1(10, 20) = %d\n", f1(10, 20))
	*/
}

// 引数, 返り値なし関数
func hello() {
	// "Hello World" を表示する
	fmt.Println("Hello World")
}

// 引数あり関数
func hello2(text string) {
	// text を表示する
	fmt.Println(text)
}

// 返り値がある関数
func add(a, b int) int {
	// a + b を計算する
	return a + b
}

// 返り値を複数返す関数
func head_tail(s []int) (int, int) {
	// sliceの先頭と末尾を返す
	return s[0], s[len(s)-1]
}

/*
func div(a, b int) (int, error) {
	// a / b を計算する
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
*/
