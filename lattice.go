package main

import (
	
)

type Lattice struct {
	chains map[string]*Blockchain
}

func (lattice *Lattice) AddChain(owner string) {

	chain := NewBlockchain(owner)
	lattice.chains[string(chain.Owner)] = chain
}

func NewBlocklattice() *Lattice {
	var genChain = GenesisChain()
	var temp map[string]*Blockchain
	temp = make(map[string]*Blockchain)
	temp[string(genChain.Owner)] = genChain
	return &Lattice{temp}
}