package main

type Blockchain struct {
	Owner  []byte
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(data string, sender string) {

	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash, sender)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockchain(owner string) *Blockchain {
	return &Blockchain{[]byte(owner), []*Block{NewGenesisBlock()}}
}

func GenesisChain() *Blockchain {
	return &Blockchain{[]byte("Genesis Chain"), []*Block{NewGenesisBlock()}}
}