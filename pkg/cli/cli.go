package cli

import (
	"awesomeProject/pkg/dao"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CLI struct {
	BlockChain *dao.BlockChain
}

func (cli *CLI) printUsage() {
	fmt.Println("用法")
	fmt.Println("addBlock 增加块")
	fmt.Println("showchain 显示区块链")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {
	cli.BlockChain.AddBlock(data)
	fmt.Println("增加成功")
}

func (cli *CLI) showBlockChain() {
	bci := cli.BlockChain.Iterator()
	for {
		block := bci.Next()
		fmt.Printf("prev hash: %x, ", block.PrevBlockHash)
		fmt.Printf("data: %s, ", block.Data)
		fmt.Printf("current hash: %x, ", block.Hash)
		pow := dao.NewProofOfWork(block)
		fmt.Printf("pow %s \n", strconv.FormatBool(pow.Validate()))
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCMD := flag.NewFlagSet("addBlock", flag.ExitOnError)
	showChainCMD := flag.NewFlagSet("showChain", flag.ExitOnError)
	addBlockData := addBlockCMD.String("data", "", "Block Data")
	switch os.Args[1] {
	case "addBlock":
		err := addBlockCMD.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "showChain":
		err := showChainCMD.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}
	if addBlockCMD.Parsed() {
		if *addBlockData == "" {
			addBlockCMD.Usage()
			os.Exit(1)
		} else {
			cli.addBlock(*addBlockData)
		}
	}
	if showChainCMD.Parsed() {
		cli.showBlockChain()
	}
}
