package dao

type BlockChain struct {
	Blocks []*Block
}

func (blocks *BlockChain) AddBlock(data string) {
	prevBlock := blocks.Blocks[len(blocks.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	blocks.Blocks = append(blocks.Blocks, newBlock)
}

func NewBlockChain() *BlockChain{
	return &BlockChain{[]*Block{NewGenesisiBlock()}}
}