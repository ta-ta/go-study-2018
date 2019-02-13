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

	// オプションのパースを利用する
	var num int
	var name string
	var help bool
	// 数字を指定できるオプションを設定する
	flag.IntVar(&num, "num", 0, "number")
	// 文字列を指定できるオプションを設定する
	flag.StringVar(&name, "name", "", "name")
	// ブール型(オプションがあるとtrue)のオプションを設定する
	flag.BoolVar(&help, "h", false, "help")
	// Parseによりos.Argsの解析を行う
	flag.Parse()

	fmt.Printf("num = %d\n", num)
	fmt.Printf("help = %v\n", help)
	fmt.Printf("args = %v\n", flag.Args())
	if help {
		flag.PrintDefaults()
	}
}
