package dao

import (
	"time"
)

type Block struct {
	TimeStamp     int64  // 19700101 00.00.00
	Data          []byte // 交易数据
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int // 工作量证明
}

//func (block *Block) SetHash() {
//	timestamp := []byte(strconv.FormatInt(block.TimeStamp, 10))
//	// 连接要哈希数据
//	headers := bytes.Join([][]byte{block.PrevBlockHash, block.Data, timestamp}, []byte{})
//	//计算哈希地址
//	hash := sha256.Sum256(headers)
//	block.Hash = hash[:]
//}

func CreateBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		TimeStamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Nonce:         0,
	}
	//block = calcu.Cal(block)
	// need run Cal(挖矿功能)

	//block.SetHash()
	return block
}

func NewGenesisiBlock() *Block {
	return CreateBlock("first", []byte{})
}
