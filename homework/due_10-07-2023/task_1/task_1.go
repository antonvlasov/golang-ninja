package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int, int) int {
	return func(x, y int) int{
		return x
	}
}

func main() {
	a := 0
	b := 1
	f := fibonacci()
	for i := 0; i < 10; i++ {
		v := f(a, b)
		fmt.Println(v)
		a, b = b, a+b
	}
}
