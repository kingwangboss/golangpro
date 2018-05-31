package main

import (
    "crypto/rand"
    "math/big"
    "fmt"
)

func main(){
    fmt.Println(RandInt64(1,80))
}

func RandInt64(min,max int64) int64{
    maxBigInt:=big.NewInt(max)
    i,_:=rand.Int(rand.Reader,maxBigInt)
    if i.Int64()<min{
        RandInt64(min,max)
    }
    return i.Int64()
}
