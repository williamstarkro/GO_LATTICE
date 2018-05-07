package main

import (
	"log"
	"fmt"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"
const accountsBucket = "accounts"

type Lattice struct {
	chains map[string]*Blockchain
	db     *bolt.DB
}

func (lattice *Lattice) AddChain(owner string) {

	err := lattice.db.Update(func(tx *bolt.Tx) error {
		a := tx.Bucket([]byte(accountsBucket))
		b := tx.Bucket([]byte(blocksBucket))

		if a == nil && b == nil {
		} else {
			if a.Get([]byte(owner)) == nil {
				chain, block := NewBlockchain(owner)
				err := b.Put(block.Hash, block.Serialize())
				if err != nil {
					log.Panic(err)
				}

				err = a.Put(chain.Owner, chain.Tip)
				if err != nil {
					log.Panic(err)
				}
				lattice.chains[string(chain.Owner)] = chain
			} else {
				fmt.Printf("%s already exists", owner)
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func NewBlocklattice() *Lattice {

	var chainsTemp map[string]*Blockchain
	chainsTemp = make(map[string]*Blockchain)
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		a := tx.Bucket([]byte(accountsBucket))
		b := tx.Bucket([]byte(blocksBucket))

		if a == nil && b == nil {
			b, err := tx.CreateBucket([]byte(blocksBucket))
			
			if err != nil {
				log.Panic(err)
			}

			a, err := tx.CreateBucket([]byte(accountsBucket))
			
			if err != nil {
				log.Panic(err)
			}

			var genChain, block = NewBlockchain("GenesisChain")
			err = b.Put(block.Hash, block.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = a.Put(genChain.Owner, genChain.Tip)
			if err != nil {
				log.Panic(err)
			}

			chainsTemp[string(genChain.Owner)] = genChain
		} else {
			a.ForEach(func(k, v []byte) error {
				bc := Blockchain{k, v}
				chainsTemp[string(k)] = &bc

				return nil
			})
			
			if err != nil {
				log.Panic(err)
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	lattice := Lattice{chainsTemp, db}

	return &lattice
}

func (lattice *Lattice) ChainIterator(owner string) *BlockchainIterator {
	return lattice.chains[owner].Iterator(lattice.db)
}