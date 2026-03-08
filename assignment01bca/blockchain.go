package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index        int
	Timestamp    string
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

var Blockchain []Block

func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	return hex.EncodeToString(hash[:])
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {

	index := len(Blockchain)
	timestamp := time.Now().String()

	data := transaction + strconv.Itoa(nonce) + previousHash + strconv.Itoa(index) + timestamp
	hash := CalculateHash(data)

	block := &Block{
		Index:        index,
		Timestamp:    timestamp,
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		Hash:         hash,
	}

	Blockchain = append(Blockchain, *block)

	return block
}

func ListBlocks() {

	fmt.Println("\n===== BLOCKCHAIN =====")

	for _, block := range Blockchain {

		fmt.Println("Index:", block.Index)
		fmt.Println("Transaction:", block.Transaction)
		fmt.Println("Nonce:", block.Nonce)
		fmt.Println("Previous Hash:", block.PreviousHash)
		fmt.Println("Current Hash:", block.Hash)
		fmt.Println("--------------------------")
	}
}

func ChangeBlock(index int, newTransaction string) {

	if index < len(Blockchain) {

		Blockchain[index].Transaction = newTransaction

		fmt.Println("Block", index, "transaction changed!")

	}
}

func VerifyChain() bool {

	for i := 1; i < len(Blockchain); i++ {

		current := Blockchain[i]
		previous := Blockchain[i-1]

		data := current.Transaction +
			strconv.Itoa(current.Nonce) +
			current.PreviousHash +
			strconv.Itoa(current.Index) +
			current.Timestamp

		recalculatedHash := CalculateHash(data)

		if current.Hash != recalculatedHash {
			fmt.Println("Hash mismatch at block", i)
			return false
		}

		if current.PreviousHash != previous.Hash {
			fmt.Println("Previous hash mismatch at block", i)
			return false
		}
	}

	return true
}