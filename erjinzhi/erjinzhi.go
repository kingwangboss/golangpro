package main

import (
	"fmt"
	"strconv"
)

func converToBin(v int) (str string) {
	result := ""
	for ; v > 0; v /= 2 {
		result = strconv.Itoa(v%2) + result
	}
	return result
}

func main() {
	fmt.Println(converToBin(21424))
}
