package main

import (
	"./BLC"
	"fmt"
)

func main() {
	block := BLC.NewBlock("Genenis Block",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	fmt.Println(block)
}