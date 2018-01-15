package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "dsjkdshdjsdh....js"
	fmt.Printf("String:%s\n,Len:%d\n,Runes:%d\n", str, len([]byte(str)), utf8.RuneCount([]byte(str)))
}
