package main

import "fmt"

func main() {
	s := "foobar"
	a := []byte(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Printf("%v\n", string(a))
}
