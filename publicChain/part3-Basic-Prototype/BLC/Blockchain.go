package BLC

type Blockchain struct {
	Blocks []*Block //存储有序的区块
}

func CreateBlockchainWithGenesisBlock() *Blockchain  {
	genesisBlock := CreateGenesisBlock("Genesis Block")
	return  &Blockchain{[]*Block{genesisBlock}}
}
