package main

import "fmt"

func fibonacci() func() int {
	prev := 0
	next := 1
	return func() int {
		temp := prev // store prev
		prev = next
		next += temp
		return temp // return stored prev
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
