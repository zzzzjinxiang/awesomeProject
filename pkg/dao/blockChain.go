package dao

import (
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "blockchain.db" //当前目录下
const blockBucket = "blocks"

type BlockChain struct {
	tip []byte
	DB  *bolt.DB
}

type BlockChainIterator struct {
	currentHash []byte
	DB          *bolt.DB
}

func (blocks *BlockChain) AddBlock(data string) *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket)) //open db
		if bucket == nil {
			gene := NewGenesisiBlock()
			bucket, err := tx.CreateBucket([]byte(blockBucket))
			if err != nil { //handle create error
				log.Panic(err)
			}

			err = bucket.Put(gene.Hash, gene.Serialize())
			if err != nil {
				log.Panic(err)
			}
			err = bucket.Put([]byte("1"), gene.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = gene.Hash
		} else {
			tip = bucket.Get([]byte("1"))
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	bc := BlockChain{
		tip: tip,
		DB:  db,
	}
	return &bc
}

func (blocks *BlockChain) Iterator() *BlockChainIterator {
	bcit := &BlockChainIterator{
		currentHash: blocks.tip,
		DB:          blocks.DB,
	}
	return bcit
}

func (it *BlockChainIterator) Next() *Block {
	var block *Block
	err := it.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		encoderBlock := bucket.Get([]byte(it.currentHash))
		block = DSerialize(encoderBlock)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	it.currentHash = block.PrevBlockHash
	return block
}

func NewBlockChain() *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket)) //open db
		if bucket == nil {
			gene := NewGenesisiBlock()
			bucket, err := tx.CreateBucket([]byte(blockBucket))
			if err != nil { //handle create error
				log.Panic(err)
			}

			err = bucket.Put(gene.Hash, gene.Serialize()) //bug 1: key required
			if err != nil {
				log.Panic(err)
			}
			err = bucket.Put([]byte("1"), gene.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = gene.Hash
		} else {
			tip = bucket.Get([]byte("1"))
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	bc := BlockChain{
		tip: tip,
		DB:  db,
	}
	return &bc
}
