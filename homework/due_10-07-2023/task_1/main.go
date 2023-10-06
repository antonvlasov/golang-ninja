package main

import "fmt"

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
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
