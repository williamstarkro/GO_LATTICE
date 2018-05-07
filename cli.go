package main

import (
	"fmt"
	"strconv"
	"flag"
	"os"
	"log"
)

type CLI struct {
	lattice *Lattice
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS")
	fmt.Println("  createwallet - Generates a new key-pair and saves it into the wallet file")
	fmt.Println("  getbalance -address ADDRESS - Get balance of ADDRESS")
	fmt.Println("  listaddresses - Lists all addresses from the wallet file")
	fmt.Println("  printchain - Print all the blocks of the blockchain")
	fmt.Println("  reindexutxo - Rebuilds the UTXO set")
	fmt.Println("  send -from FROM -to TO -amount AMOUNT -mine - Send AMOUNT of coins from FROM address to TO. Mine on the same node, when -mine is set.")
	fmt.Println("  startnode -miner ADDRESS - Start a node with ID specified in NODE_ID env. var. -miner enables mining")
}

func (cli *CLI) newAccount(account string) {
	cli.lattice.AddChain(account)
	fmt.Println("Success!")
}

func (cli *CLI) addBlock(data, account string) {
	cli.lattice.chains[account].AddBlock(data, account, cli.lattice.db)
	fmt.Println("Success!")
}

func (cli *CLI) printChain(account string) {
	bci := cli.lattice.ChainIterator(account)

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

func (cli *CLI) Run() {
	//cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	newAccountCmd := flag.NewFlagSet("newaccount", flag.ExitOnError)
	newAccountData := newAccountCmd.String("account", "", "New Account Name")
	printChainData := printChainCmd.String("account", "", "Account Chain")
	addBlockAccount := addBlockCmd.String("account", "", "Account Chain")
	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "newaccount":
		err := newAccountCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" ||  *addBlockAccount == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockAccount, *addBlockData)
	}

	if printChainCmd.Parsed() {
		if *printChainData == "" {
			printChainCmd.Usage()
			os.Exit(1)
		}
		cli.printChain(*printChainData)
	}

	if newAccountCmd.Parsed() {
		if *newAccountData == "" {
			newAccountCmd.Usage()
			os.Exit(1)
		}
		cli.newAccount(*newAccountData)
	}
}