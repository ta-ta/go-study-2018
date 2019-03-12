package main

import "fmt"

// 関数FizzBuzz() を作成する
// - 引数: int n
// - 返り値なし
// - 1 から n までの整数i に対して以下を行う
// 	- i が 3 の倍数の時, "Fizz" を表示
// 	- i が 5 の倍数の時, "Buzz" を表示
// 	- i が 3 の倍数 かつ 5 の倍数の時, "FizzBuzz" を表示
// 	- それ以外の時, i を表示

// FizzBuzz()に修正を加える
// - i が 3 のつく数字のときも "Fizz" を表示する
// - i が 5 のつく数字のときも "Buzz" を表示する

func main() {
	FizzBuzz(100)
}

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
