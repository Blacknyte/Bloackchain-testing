package main

import (
	"fmt"
	"assignment01bca"
)

func main() {

	// Genesis Block
	genesis := assignment01bca.NewBlock(
		"Genesis Block - 22i-1692",
		18,
		"0",
	)

	// Block 1
	block1 := assignment01bca.NewBlock(
		"alice pays bob 5 coins - 692",
		5,
		genesis.Hash,
	)

	// Block 2
	assignment01bca.NewBlock(
		"bob pays charlie 2 coins",
		3,
		block1.Hash,
	)

	// Print blockchain
	assignment01bca.ListBlocks()

	// Verify chain
	fmt.Println("\nVerifying Blockchain...")
	if assignment01bca.VerifyChain() {
		fmt.Println("Blockchain is VALID")
	} else {
		fmt.Println("Blockchain is INVALID")
	}

	// Tamper with Block 1
	fmt.Println("\nTampering Block 1...")
	assignment01bca.ChangeBlock(1, "Hacker changed transaction")

	// Verify again
	fmt.Println("\nVerifying Blockchain Again...")
	if assignment01bca.VerifyChain() {
		fmt.Println("Blockchain is VALID")
	} else {
		fmt.Println("Blockchain is INVALID")
	}
}