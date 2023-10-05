package main

import "fmt"
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    a, b := 1, -1 
    return func() int{
        a, b = a + b, a
        return a 
    }
}

func main() {
	f := fibonacci()
    printf := fmt.Printf

	for i := 0; i < 10; i++ {
		printf("%d\n", f())
	}
}

