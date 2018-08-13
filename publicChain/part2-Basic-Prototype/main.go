package main

import (
	"./BLC"
	"fmt"
)

func main() {
	genesisblock := BLC.CreateGenesisBlock("Genesis Block")
	fmt.Println(genesisblock)
}