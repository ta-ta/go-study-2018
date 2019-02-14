package main

//
// JSON Formatterの仕様
// - コマンドライン引数で与えられたファイルの内容をJSONとして解釈,
//   pretty printとして出力
// - コマンドラインオプションでインデントレベル(タブの個数)を指定可能
//
// 追加仕様
// - コマンドライン引数で与えられたファイルの内容をJSONとして解釈,
//   minifyして出力する機能を追加
// - pretty print, minifyのどちらで出力するかは,
//   コマンドラインオプションの'mode'で変更可能
//

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type JsonFormatter interface {
	Format() (string, error)
}

type JsonPrettyPrinter struct {
	IndentLevel int
	Data        []byte
}

func (printer JsonPrettyPrinter) Format() (string, error) {
	buf := bytes.Buffer{}
	indent := strings.Repeat("\t", printer.IndentLevel)
	err := json.Indent(&buf, printer.Data, "", indent)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

type JsonMinifyPrinter struct {
	Data []byte
}

func (printer JsonMinifyPrinter) Format() (string, error) {
	buf := bytes.Buffer{}
	err := json.Compact(&buf, printer.Data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func main() {
	var mode string
	var level int
	var fileName string
	flag.StringVar(&mode, "mode", "pp", "mode")
	flag.IntVar(&level, "level", 1, "indent level")
	flag.StringVar(&fileName, "file", "", "input file")
	flag.Parse()

	var printer JsonFormatter
	var inputFile *os.File
	if fileName == "" {
		inputFile = os.Stdin
	} else {
		file, err := os.Open(fileName)
		if err != nil {
			log.Println(err)
			return
		}
		inputFile = file
	}
	data, err := ioutil.ReadAll(inputFile)
	if err != nil {
		log.Println(err)
	}
	if inputFile != os.Stdin {
		inputFile.Close()
	}

	switch mode {
	case "mini":
		printer = JsonMinifyPrinter{Data: data}
	default:
		printer = JsonPrettyPrinter{Data: data, IndentLevel: level}
	}
	output, err := printer.Format()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%s\n", string(output))
}
