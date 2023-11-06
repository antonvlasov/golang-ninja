package main

import "fmt"

<<<<<<< HEAD
func fibonacci() func() int {
	prev := 0
	next := 1
	return func() int {
		temp := prev // store prev
		prev = next
		next += temp
		return temp // return stored prev
	}
=======
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(k int) int {
	fib := func(k int) int {
		i, i1, cnt, ans := 1, 1, 1, 1
		if k <= 1 {
			return k
		}
		for cnt <= k/2 {
			ans = i
			i = i * (2*i1 - i)
			i1 = i1*i1 + ans*ans
			cnt *= 2
		}
		if cnt == k {
			return i
		} else if cnt+1 == k {
			return i1
		}
		cnt += 2
		ans = i1 + i
		for ; cnt < k; cnt++ {
			ans += i1
			i1 = ans - i1
		}
		return ans
	}
	return fib
>>>>>>> upstream/main
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
<<<<<<< HEAD
		fmt.Println(f())
=======
		fmt.Println(f(i))
>>>>>>> upstream/main
	}
}
