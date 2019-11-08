package main

import (
	bcs "awesomeProject/pkg/dao"
	"awesomeProject/pkg/funcs"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("begin")
	bc := bcs.NewBlockChain()
	bc.AddBlock("first trial1")
	bc.AddBlock("first trial2")
	bc.AddBlock("first trial3")

	for _, block := range bc.Blocks {
		block = funcs.Cal(block)
		fmt.Printf("prev hash: %x, ", block.PrevBlockHash)
		fmt.Printf("data: %s, ", block.Data)
		fmt.Printf("current hash: %x, ", block.Hash)
		pow := funcs.NewProofOfWork(block)
		fmt.Printf("pow %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}
