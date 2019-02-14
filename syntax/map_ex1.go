package main

import "fmt"

func main() {
	// 課題: 変数strの文字列の中に、各文字が何回現れるかを出力する
	str := "abracadabra"

	// Hint: string型の要素を参照した時の型: rune型
	// str[0] => 'a' (rune型)

	counterMap := map[rune]int{}
	for _, c := range str {
		counterMap[c] += 1
	}

	for k, v := range counterMap {
		fmt.Printf("%c: %d\n", k, v)
	}
}
