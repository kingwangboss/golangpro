package main

import "fmt"

func main() {
	var arr [10]int
	for index := 0; index < 10; index++ {
		arr[index] = index
	}
	fmt.Printf("%v", arr)
}
