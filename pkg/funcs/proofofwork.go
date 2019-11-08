package funcs

import (
	"awesomeProject/pkg/dao"
	"awesomeProject/pkg/utils"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

const targetBits = 24

type ProofOfWork struct {
	block  *dao.Block
	target *big.Int
}

func NewProofOfWork(block *dao.Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{block, target}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			utils.IntToHex(pow.block.TimeStamp),
			utils.IntToHex(int64(targetBits)),
			utils.IntToHex(int64(nonce))},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("当前区块数据%s", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x\n", hash)
		//fmt.Printf("%d\n", nonce)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 { // 挖矿校验
			break
		} else {
			nonce++
		}
	}
	fmt.Println("\n\n")
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {

	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isVaild := (hashInt.Cmp(pow.target) == -1)
	return isVaild
}
