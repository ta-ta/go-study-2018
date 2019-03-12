package main

// JSON Formatter
// コマンドラインオプションで指定したファイルをJSONとして解釈し,
// pretty printで出力

// コマンドラインオプション`mode`でpretty print, minifyを切り替えられるようにする
// コマンドラインオプション`level`でインデントレベル(タブの個数)を指定できるようにする

// 構造体を使うなど実装を工夫してみてください

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

type JsonFormatter struct {
	Mode        string
	IndentLevel int
	Data        []byte
}

func (printer JsonFormatter) Format() (string, error) {
	buf := bytes.Buffer{} // Buffer型の変数を用意

	if printer.Mode == "mini" {
		err := json.Compact(&buf, printer.Data) // minifyしてbufに格納
		if err != nil {
			return "", err
		}
	} else {
		indent := strings.Repeat("\t", printer.IndentLevel) // 文字列を繰り返す
		err := json.Indent(&buf, printer.Data, "", indent)  // indentされた文字列をbufに格納
		if err != nil {
			return "", err
		}
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

	printer = JsonFormatter{Data: data, Mode: mode, IndentLevel: level}

	output, err := printer.Format()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%s\n", string(output))
}
