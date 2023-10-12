package main

import "fmt"

func fibonacci(a, b int) (int, int) {
	a, b = b, a+b
	return a, b
}

func main() {
	a := 0
	b := 1
	fmt.Println(0)
	for i := 0; i < 10; i++ {
		a, b = fibonacci(a, b)
		fmt.Println(a)
	}
}
