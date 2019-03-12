package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// コマンドライン引数をそのまま利用する
	fmt.Printf("os.Args = %v\n", os.Args)
	for i, a := range os.Args {
		// コマンドライン引数の型は文字列
		fmt.Printf("os.Args[%d] = '%s'\n", i, a)
	}

	// `go run arg.go -num 123` のように実行した時,
	// numという名前のオプションとして123を設定したい

	// オプションのパースを利用する
	var num int
	var name string
	var help bool

	// 数字を指定できるオプションを設定する
	flag.IntVar(&num, "num", 0, "number")
	// 第一引数: 解析した値をいれておく変数
	// 第二引数: 実行時にこのオプションを指定するための文字列
	// 第三引数: 指定しなかった場合の値
	// 第四引数: flag.PrintDefaults() により表示される説明分
	// 文字列を指定できるオプションを設定する
	flag.StringVar(&name, "name", "", "name")
	// ブール型のオプションを設定する (オプションを指定したときtrue)
	flag.BoolVar(&help, "h", false, "help")

	// Parseによりos.Argsの解析を行う
	flag.Parse()

	fmt.Printf("num = %d\n", num)
	fmt.Printf("name = %v\n", name)
	fmt.Printf("help = %v\n", help)

	fmt.Printf("args = %v\n", flag.Args()) // オプション以外の引数

	if help {
		flag.PrintDefaults()
	}
}
