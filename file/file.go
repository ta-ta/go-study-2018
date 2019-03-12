package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	ReadFileLine("file/test.txt")

	/*
		ReadFile("file/test.txt")
	*/
}

// ReadFileLine ファイルの内容を一行ずつ読み込む
func ReadFileLine(filename string) {
	// ファイルを開く
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ファイルの内容を一行ずつ読む
	reader := bufio.NewReader(file) //Reader を返す
	i := 1
	for {
		line, _, err := reader.ReadLine() // ReadLine によって一行ずつ, byte配列を返す
		if err == io.EOF {
			break
		}
		fmt.Printf("%d行目: %s\n", i, line)
		i++
	}

	// ファイルを閉じる
	file.Close()
}

/*
// ReadFile ファイルの内容を一括で読み込む
func ReadFile(filename string) {
	// ファイルを開く
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ファイルの内容全てを読む
	data, err := ioutil.ReadAll(file) // byte配列を返す
	if err != nil {
		fmt.Println(err)
		return
	}

	// 内容を出力
	fmt.Printf("===== %s\n", filename)
	fmt.Printf("%s\n", data) // []byteも%s
	fmt.Println("=====")

	// ファイルを閉じる
	file.Close()
}
*/
