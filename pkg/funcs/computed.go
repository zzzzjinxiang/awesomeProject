package funcs

import (
	"awesomeProject/pkg/dao"
)

func Cal(block *dao.Block) *dao.Block {
	pow := dao.NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}
