package main

import "fmt"

func main() {
	str := "A"
	for i := 0; i < 100; i++ {
		fmt.Printf("%v\n", str)
		str = str + "A"
	}
}
