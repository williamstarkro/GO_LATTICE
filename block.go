package main

import (
	"bytes"
	"strconv"
	"crypto/sha256"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	Sender        []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {

	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Sender, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte, sender string) *Block {

	block := &Block{time.Now().Unix(), []byte(data), []byte(sender), prevBlockHash, []byte{}}
	block.SetHash()

	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", "Genesis", []byte{})
}