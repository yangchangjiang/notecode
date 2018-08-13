package BLC

import "math/big"

//0000 0000 0000 0000 1001 0001 0000 .... 0001
//256位Hash里面前面至少要有16个0
const targetBit  = 16

type ProofOfWork struct {
	Block *Block //当前要验证的区块
	//diff int64
	target *big.Int //大数据存储
}

func (proofOfWork *ProofOfWork) Run() ([]byte,int64) {
	return nil,0
}

//创建新的工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {
	//1、创建一个初始值为1的target

	target := big.NewInt(1)
	//2、左移动256-targetBit
	target = target.Lsh(target,256 - targetBit)

	return &ProofOfWork{block,target}
}