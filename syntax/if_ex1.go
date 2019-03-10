package main

import "fmt"

func main() {
	// 以下の条件式を修正し v が奇数の時は"奇数", 偶数の時は"偶数"と表示されるように実装
	v := 123
	if v%2 == 0 { // 修正箇所
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}
