package main

import "fmt"

func fibomacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 20)
	go fibomacci(cap(c), c)
	for v := range c {
		fmt.Printf("%v ", v)
	}
}
