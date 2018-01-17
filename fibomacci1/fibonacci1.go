package main

import "fmt"

func fibomacci1(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for index := 0; index < 10; index++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibomacci1(c, quit)
}
