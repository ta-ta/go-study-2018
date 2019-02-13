package main

import "fmt"

// タスク: FizzBuzz

// nを受け取り、1~nのFizzBuzzを出力する関数FizzBuzz
func main() {
	FizzBuzz(100)
}

// ** 回答1
func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}

// ** 回答2
func FizzBuzz2(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Printf("%d\n", i)
		}
	}
}
