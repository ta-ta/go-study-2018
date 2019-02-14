package main

import (
	"fmt"
	"math"
)

// 課題: 各構造体にフィールドとArea()メソッドを追加して面積を計算できるようにする

// Shapeインタフェース
type Shape interface {
	// 面積を計算するArea()メソッド
	Area() float64
}

// Square構造体
type Square struct {
	X float64
}

func (s Square) Area() float64 {
	return s.X * s.X
}

// Rectangle構造体
type Rectangle struct {
	A, B float64
}

func (r Rectangle) Area() float64 {
	return r.A * r.B
}

// Triangle構造体
type Triangle struct {
	A, B, Rad float64
}

func (t Triangle) Area() float64 {
	return t.A * t.B * math.Sin(t.Rad) / 2
}

// 1辺の長さxを入力として初期化する
func NewSquare(x float64) *Square {
	return &Square{X: x}
}

// 2辺の長さを入力として初期化する
func NewRectangle(a, b float64) *Rectangle {
	return &Rectangle{A: a, B: b}
}

// 2辺(a, b)とその間の角(rad)(ラジアン)を入力として初期化する
func NewTriangle(a, b, rad float64) *Triangle {
	return &Triangle{A: a, B: b, Rad: rad}
}

func main() {
	var shapes []Shape
	shapes = append(shapes, NewSquare(4))
	shapes = append(shapes, NewRectangle(4, 5))
	shapes = append(shapes, NewTriangle(2, 4, math.Pi/6))

	for _, s := range shapes {
		fmt.Printf("Area() = %f\n", s.Area())
	}
}
