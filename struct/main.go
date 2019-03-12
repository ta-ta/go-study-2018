package main

import "fmt"

// 構造体の定義
type Pair struct {
	// フィールドのリストを書く
	Left  int
	Right int
}

// メソッドの定義
func (p Pair) String() string {
	return fmt.Sprintf("(%d, %d)", p.Left, p.Right)
}

func (p Pair) Sum() int {
	return p.Left + p.Right
}

func main() {
	p := Pair{Left: 10, Right: 20}
	// 構造体の中身を表示
	fmt.Printf("%#v\n", p)
	fmt.Println()

	// メソッド呼び出し
	fmt.Printf("p.String() = %s\n", p.String())
	fmt.Printf("p.Sum() = %d\n", p.Sum())
	fmt.Println()

	/*
		p.Double1()
		fmt.Printf("p.Double1() => %#v\n\n", p)

		p.Double2()
		fmt.Printf("p.Double2() => %#v\n\n", p)


		// インターフェスの変数を宣言
		var printer Printer

		// Pairを代入
		printer = p
		printer.Print()

		// DummyPrinterを代入
		printer = DummyPrinter{}
		printer.Print()
	*/
}

/*
// 構造体の値を変更する(NG)
func (p Pair) Double1() {
	p.Left *= 2
	p.Right *= 2
}

// 構造体の値を変更する(OK)
func (p *Pair) Double2() {
	p.Left *= 2
	p.Right *= 2
}

// インタフェースの定義
type Printer interface {
	Print()
}

func (p Pair) Print() {
	fmt.Println(p.String())
}

// Pairとは別の構造体DummyPrinter
type DummyPrinter struct {
}

func (p DummyPrinter) Print() {
	fmt.Println("dummy printer")
}
*/
