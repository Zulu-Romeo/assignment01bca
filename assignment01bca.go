//Zein Rohail
//20I-2441
//Assignmet 1
//Blockchain & Cyrptocurrency

package assignment01bca

import (
	"crypto/sha256"
	"fmt"
)

// Block represents a single block in the blockchain.
type Block struct {
	Transaction  string
	Nonce        int
	Hash         string
	PreviousHash string
}

type Blockchain struct {
	Chain []Block
}

// Creates a new block to the blockchain.
func (obj *Blockchain) NewBlock(transaction string, nonce int, previousHash string) {
	block := Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}

	block.Hash = obj.CreateHash(fmt.Sprintf("%v%v%v", block.Transaction, block.Nonce, block.PreviousHash))

	obj.Chain = append(obj.Chain, block)
}

// Prints List of Blocks
func (obj *Blockchain) ListBlocks() {
	for i, block := range obj.Chain {
		fmt.Printf("============================================\n")
		fmt.Printf("Block No. # %d:\n", i+1)
		fmt.Printf("  Transaction: %s\n", block.Transaction)
		fmt.Printf("  Nonce: %d\n", block.Nonce)
		fmt.Printf("  Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("  Current Hash: %s\n\n", block.Hash)
	}
}

// Malicious Transaction
func (obj *Blockchain) ChangeBlock(blockIndex int, newTransaction string) {
	if blockIndex >= 0 && blockIndex < len(obj.Chain) {
		obj.Chain[blockIndex].Transaction = newTransaction
		obj.Chain[blockIndex].Hash = obj.CreateHash(fmt.Sprintf("%v%v%v", obj.Chain[blockIndex].Transaction, obj.Chain[blockIndex].Nonce, obj.Chain[blockIndex].PreviousHash))
	}
}

// Comparing hashes to see validity.
func (obj *Blockchain) VerifyChain() bool {
	for i := 1; i < len(obj.Chain); i++ {
		currentBlock := obj.Chain[i]
		previousBlock := obj.Chain[i-1]

		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}

	}
	return true
}

// Calculate Hashes
func (obj *Blockchain) CreateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	return fmt.Sprintf("%x", hash)
}
