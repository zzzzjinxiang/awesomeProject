package dao

import (
	"bytes"
	"encoding/gob"
	"log"
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
	block = runfunc(block)
	//block.SetHash()
	return block
}

func runfunc(block *Block) *Block {
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func NewGenesisiBlock() *Block {
	return CreateBlock("first", []byte{})
}

// 对象转化为二进制，写入文件
func (block *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}

	return res.Bytes()
}

//读取文件，二进制转对象
func DSerialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
