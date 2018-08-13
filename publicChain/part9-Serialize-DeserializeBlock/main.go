package main

import (
	"./BLC"
	"fmt"
)
func main() {
	////创世区块
	//blockchain := BLC.CreateBlockchainWithGenesisBlock()
	//
	////新区块
	//blockchain.AddBlockToBlockchain("Send 100 RMB to CJ1",blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//blockchain.AddBlockToBlockchain("Send 300 RMB to CJx",blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//blockchain.AddBlockToBlockchain("Send 400 RMB to CJz",blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//blockchain.AddBlockToBlockchain("Send 200 RMB to CJf",blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//fmt.Println(blockchain)
	//fmt.Println(blockchain.Blocks)
	block := BLC.NewBlock("Test",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	fmt.Printf("%d\n",block.Nonce)
	fmt.Printf("%d\n",block.Hash)

	proofOfWork := BLC.NewProofOfWork(block)
	fmt.Printf("%v",proofOfWork.IsValid())

	bytes := block.Serialize()
	fmt.Println(bytes)

	block = BLC.DeserializeBlock(bytes)
	fmt.Printf("%d\n",block.Nonce)
	fmt.Printf("%d\n",block.Hash)


}