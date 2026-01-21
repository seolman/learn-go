package main

func concurrentFib(n int) []int {
	ch := make(chan int)
	go fibonacci(n, ch)
	s := make([]int, n)
	for i := range n {
		num := <-ch
		s[i] = num
	}
	return s
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for range n {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
