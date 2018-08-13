package main

import (
	"./BLC"
	"fmt"
)

func main() {
	//创世区块
	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	fmt.Println(blockchain)
	fmt.Println(blockchain.Blocks)
	fmt.Println(blockchain.Blocks[0])

	//新区块
	blockchain.AddBlockToBlockchain("Send 100 RMB to CJ1",blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("Send 300 RMB to CJx",blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("Send 400 RMB to CJz",blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("Send 200 RMB to CJf",blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	fmt.Println(blockchain)

	fmt.Println(blockchain.Blocks[0])
	fmt.Println(blockchain.Blocks[1])
	fmt.Println(blockchain.Blocks[2])
	fmt.Println(blockchain.Blocks[3])
	fmt.Println(blockchain.Blocks[4])
	
}