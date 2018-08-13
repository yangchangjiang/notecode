package BLC

import (
	"time"
	"fmt"
)

type Block struct {
	//1、区块高度
	Height int64
	//2、上一个区块的HASH
	PrevBlockHash []byte
	//3、交易数据
	Data []byte
	//4、时间戳
	Timstamp int64
	//5、Hash
	Hash [] byte

	Nonce int64
}


// 1、创建新区块
func NewBlock(data string,height int64,prevBlockHash []byte) *Block {
	//创建区块
	block := &Block{height,prevBlockHash,[]byte(data),time.Now().Unix(),nil,0}
	//调用工作量证明的方法并且返回有效的Hash和Nonce值

	pow := NewProofOfWork(block)
	//00000
	hash,nonce := pow.Run()
	block.Hash=hash
	block.Nonce=nonce

	fmt.Println()
	return block
}
// 2、单独方法，生成创世区块
func CreateGenesisBlock(data string)  *Block{
	return NewBlock(data,1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
}
