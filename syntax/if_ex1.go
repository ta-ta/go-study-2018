package main

import "fmt"

func main() {
	v := 123
	// v が奇数の時は"奇数", 偶数の時は"偶数"と表示されるように実装
	if v%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}
