package main

import "fmt"

// 課題1:
// 関数FizzBuzz() を作成する
// - 引数: 自然数n
// - 返り値なし
// - 1からnまで順に以下を行う
// 	- nが3の倍数の時, "Fizz"を表示
// 	- nが5の倍数の時, "Buzz"を表示
// 	- nが3の倍数 かつ 5の倍数の時, "FizzBuzz"を表示
// 	- それ以外の時, n を表示
//
// // 課題2:
// FizzBuzz()に修正を加える
// - nが3のつく数字のときも"Fizz"を表示する
// - nが5のつく数字のときも"Buzz"を表示する

func main() {
	N := 100
	FizzBuzz(N)
}

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%15 == 0 || (check(i, 3) && check(i, 5)) || (i%3 == 0 && check(i, 5)) || (i%5 == 0 && check(i, 3)) {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 || check(i, 3) {
			fmt.Println("Fizz")
		} else if i%5 == 0 || check(i, 5) {
			fmt.Println("Buzz")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}

func check(n, m int) bool {
	// nがmのつく数字ならtrue
	for n > 0 {
		if n%10 == m {
			return true
		}
		n = n / 10
	}
	return false
}
