package main

type Blockchain struct {
	owner  []byte
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string, sender string) {

	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, sender, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewBlockchain(owner string) *Blockchain {
	return &Blockchain{[]byte(owner), []*Block{NewGenesisBlock()}}
}

func GenesisChain() *Blockchain {
	return &Blockchain{[]byte("Genesis Chain"), []*Block{NewGenesisBlock()}}
}