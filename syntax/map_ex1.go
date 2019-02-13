package main

import "fmt"

func main() {
	str := "abracadabra"
	// Hint: string型の要素を参照した時の型: rune型
	// str[0] => rune型

	counterMap := map[rune]int{}
	for _, c := range str {
		counterMap[c] += 1
	}

	for k, v := range counterMap {
		fmt.Printf("%c: %d\n", k, v)
	}
}
