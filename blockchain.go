package main

import (
	"log"

	"github.com/boltdb/bolt"
)

type Blockchain struct {
	Owner  []byte
	Tip    []byte
}

func (bc *Blockchain) AddBlock(data, sender string, db *bolt.DB) {
	var lastHash []byte

	err := db.View(func(tx *bolt.Tx) error {
		a := tx.Bucket([]byte(accountsBucket))
		lastHash = a.Get([]byte(bc.Owner))

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, lastHash, sender)

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		a := tx.Bucket([]byte(accountsBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = a.Put(bc.Owner, newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		bc.Tip = newBlock.Hash

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func (bc *Blockchain) Iterator(db *bolt.DB) *BlockchainIterator {
	bci := &BlockchainIterator{bc.Tip, db}

	return bci
}

func NewBlockchain(owner string) (*Blockchain, *Block) {
	genesis := NewGenesisBlock()
	return &Blockchain{[]byte(owner), genesis.Hash}, genesis
}