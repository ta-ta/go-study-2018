package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	headLine := Head("test.txt")
	fmt.Printf("head = %s\n", headLine)
	tailLine := Tail("test.txt")
	fmt.Printf("tail = %s\n", tailLine)
}

// ファイルの先頭行を表示する
func Head(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	line, _, err := reader.ReadLine()
	if err != nil {
		return ""
	}
	return string(line)
}

// ファイルの最終行を表示する
func Tail(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	var last string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		last = string(line)
	}
	return last
}
