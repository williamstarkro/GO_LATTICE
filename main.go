package main

import (
	"fmt"
)

func main() {
	lattice := NewBlocklattice()
	fmt.Printf("Gen name: %s\n", lattice.chains["Genesis Chain"].Owner)
	lattice.AddChain("Will")
	lattice.chains["Will"].AddBlock("Send 1 to Willy", "Genesis")
	lattice.chains["Will"].AddBlock("Send 2 more to Willy", "Genesis")

	for _, block := range lattice.chains["Will"].Blocks {
		fmt.Printf("Prev block hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Sender: %s\n", block.Sender)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Time: %d\n", block.Timestamp)
		fmt.Println()
	}
}