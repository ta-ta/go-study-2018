package main

import "fmt"

func main() {
	// 連想配列: mapの宣言
	// keyの値はint型, valueの値はstring型
	m1 := map[int]string{}

	m1[10] = "abc"
	m1[20] = "cde"

	// mapの中身確認
	fmt.Printf("%v\n", m1)
	fmt.Println()
	fmt.Println()

	// mapの要素へのアクセス
	fmt.Printf("m1[10] = %s, m1[20] = %s\n", m1[10], m1[20])

	// 存在しない要素へのアクセス: 0値が返る
	fmt.Printf("m1[30] = %s\n\n", m1[30])
	fmt.Println()

	// 要素が存在するかの判定
	// 第二引数を指定すると存在するかのbool値が返る
	// 要素が存在する: ok = true
	// 要素が存在しない: ok = false
	if el, ok := m1[10]; ok {
		fmt.Printf("m1[10]は存在 (value = %s)\n\n", el)
	}
	if el, ok := m1[30]; ok {
		fmt.Printf("m1[30]は存在 (value = %s)\n\n", el)
	}
	fmt.Println()
	fmt.Println()

	// 要素を含めた宣言
	m2 := map[string]int{
		"abc": 123,
		"cde": 456,
	}

	fmt.Printf("m2['abc'] = %d, m2['cde'] = %d\n\n", m2["abc"], m2["cde"])

	// for文でのmapの参照
	// 要素の順序はランダムになることに注意
	for key, value := range m2 {
		fmt.Printf("key = %s, value = %d\n", key, value)
	}
}
