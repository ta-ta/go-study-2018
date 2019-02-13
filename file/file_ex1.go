package main

import (
	"fmt"
)

func main() {
	headLine := Head("test.txt")
	fmt.Printf("head = %s\n", headLine)
	tailLine := Tail("test.txt")
	fmt.Printf("tail = %s\n", tailLine)
}

// ファイルの先頭行を表示する
func Head(filename string) string {
	return ""
}

// ファイルの最終行を表示する
func Tail(filename string) string {
	return ""
}
