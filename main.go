package main

import (
	"fmt"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Willy")
	bc.AddBlock("Send 2 more to Willy")

	for _, block := range bc.blocks {
		fmt.Printf("Prev block hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Time: %d\n", block.Timestamp)
		fmt.Println()
	}
}