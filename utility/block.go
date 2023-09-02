// contains the block struct and methods to create and manipulate blocks

package utility

import (
	"time"
)

// Block represents each 'item' in the blockchain
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nounce        int
}

// Blockchain represents a list of blocks
type Blockchain struct {
	BlockArray []*Block
}

// creating a new block
func NewBlock(data string, prevBlockHash []byte) (*Block, error) {
	block := &Block{Timestamp: time.Now().Unix(), Data: []byte(data), PrevBlockHash: prevBlockHash}

	pow := NewProofOfWork(block)
	nonce, hash, err := pow.Run()
	if err != nil {
		return nil, err
	}

	block.Nounce = nonce
	block.Hash = hash[:]
	return block, nil
}

// creating the genesis block
func NewGenesisBlock() *Block {
	new_block, _ := NewBlock("Genesis Block", []byte{})
	return new_block
}

// creating a new blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// adding a block to the blockchain
func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.BlockArray[len(chain.BlockArray)-1]
	newBlock, err := NewBlock(data, prevBlock.Hash)
	if err != nil {
		panic(err)
	}
	chain.BlockArray = append(chain.BlockArray, newBlock)
}
