package BLC

type Blockchain struct {
	Blocks []*Block //存储有序的区块
}

// 增加区块到区块链里面
func (blc *Blockchain)AddBlockToBlockchain(data string,height int64,preHash []byte)  {
	//创建新区快
	newBlock := NewBlock(data,height,preHash)
	//往链中添加区块
	blc.Blocks = append(blc.Blocks,newBlock)
}

func CreateBlockchainWithGenesisBlock() *Blockchain  {
	genesisBlock := CreateGenesisBlock("Genesis Block")
	return  &Blockchain{[]*Block{genesisBlock}}
}
