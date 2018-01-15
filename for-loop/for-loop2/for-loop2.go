package main

import "fmt"

func main() {
	i := 1

I:
	fmt.Println(i)

	if i < 10 {
		i++
		goto I
	}
}
