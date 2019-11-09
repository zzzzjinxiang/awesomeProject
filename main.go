package main

import (
	cli2 "awesomeProject/pkg/cli"
	"awesomeProject/pkg/dao"
)

func main() {
	block := dao.NewBlockChain()
	defer block.DB.Close()
	cli := cli2.CLI{block}
	cli.Run()

}
