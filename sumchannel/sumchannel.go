package main

import "fmt"

func sum(arr []int, c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum
}

func main() {
	arr := []int{3, 23, 5, 6, 32, 3, 7, 2, 4, 3, 73, 4, 7, 64, 3,
		2, 32, 3, 23, 5, 6, 32, 3, 7, 2, 4, 3, 73, 4, 7, 64, 3, 2,
		32, 3, 23, 5, 6, 32, 3, 7, 2, 4, 3, 73, 4, 7, 64, 3, 2, 32,
		32, 3, 23, 5, 6, 32, 3, 7, 2, 4, 3, 73, 4, 7, 64, 3, 2, 32,
		32, 3, 23, 5, 6, 32, 3, 7, 2, 4, 3, 73, 4, 7, 64, 3, 2, 32,
		32, 3, 23, 5, 6, 32, 3, 7, 2, 4, 3, 73, 4, 7, 64, 3, 2, 32,
		32, 3, 23, 5, 6, 32, 3, 7, 2, 4, 3, 73, 4, 7, 64, 3, 2, 32,
		32, 3, 23, 5, 6, 32, 3, 7, 2, 4, 3, 73, 4, 7, 64, 3, 2, 32,
		32, 3, 23, 5, 6, 32, 3, 7, 2, 4, 3, 73, 4, 7, 64, 3, 2, 32}
	c := make(chan int)
	go sum(arr[:len(arr)/2], c)
	go sum(arr[len(arr)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
