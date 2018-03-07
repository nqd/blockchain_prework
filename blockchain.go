package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// Blockchain is our global blockchain.
var Blockchain []Block

// Block is our basic data structure!
type Block struct {
	Data      string
	Timestamp int64
	PrevHash  []byte
	Hash      []byte
}

// InitBlockchain creates our first Genesis node.
func InitBlockchain() {
	bc := Block{
		"Genesis Block",
		time.Now().Unix(),
		[]byte{},
		[]byte{},
	}
	bc.Hash = bc.calculateHash()
	Blockchain = []Block{bc}
	spew.Dump(Blockchain)
}

// NewBlock creates a new Blockchain Block.
func NewBlock(oldBlock Block, data string) Block {
	bc := Block{
		data,
		time.Now().Unix(),
		oldBlock.Hash,
		[]byte{},
	}
	bc.Hash = bc.calculateHash()
	return bc
}

// AddBlock adds a new block to the Blockchain.
func AddBlock(b Block) error {
	fmt.Println("******TODO: IMPLEMENT AddBlock!******")
	spew.Dump(Blockchain)
	// Fill me in, brave wizard.
	return nil
}

func (b *Block) calculateHash() []byte {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	data := []byte(b.Data)
	headers := bytes.Join([][]byte{b.PrevHash, data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	return hash[:]
}
