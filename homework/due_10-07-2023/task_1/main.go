package main

import "fmt"

func fib(a int) int {
	switch {
	case a == 0:
		return 0
	case a == 1:
		return 1
	default:
		return fib(a-1) + fib(a-2)
	}

}
func main() {
	fmt.Println(fib(0))
}
