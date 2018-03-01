package main

import (
	"fmt"
	. "github.com/jeffotoni/blockchain/src/blockchain"
)

func main() {

	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Walker")
	bc.AddBlock("Send 2 BTC to Alencar")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
