package chain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"owlchain.localhost/block"
	"time"
)

type Chain struct {
	blocks       []block.Block
	workingBlock block.Block
	difficulty   int // num of leading 0 bytes in hash required
}

func NewChain(difficulty int) *Chain {
	ch := Chain{
		blocks:       make([]block.Block, 0),
		workingBlock: *block.NewBlock(),
		difficulty:   difficulty,
	}
	return &ch
}

func (ch *Chain) AddWorkingBlock() {
	ch.blocks = append(ch.blocks, ch.workingBlock)
}

// HashWorkingBlock finds a valid hash for the working block
func (ch *Chain) HashWorkingBlock() ([]byte, uint64) {
	compBytes := make([]byte, ch.difficulty)

	// serialize block for first time
	sr := make([]byte, block.LengthBytes)
	copy(sr, ch.workingBlock.Serialize(false))
	currNonce := uint64(0)

	for {
		// hash the block
		hx := sha256.New()
		hx.Write(sr)
		digest := hx.Sum(nil)

		// check for a success
		if bytes.Equal(digest[0:ch.difficulty], compBytes) {
			ch.workingBlock.UpdateHash(digest, currNonce)
			return digest, currNonce
		}

		// bad hash, update nonce, try again!
		currNonce++
		binary.PutUvarint(sr[block.HashLengthBytes*2:], currNonce)
	}
}

func (ch *Chain) ResetWorkingBlock(prev []byte) {
	bk := block.NewBlockWithPrev(prev)
	ch.workingBlock = *bk
}

func (ch *Chain) Run() {
	for {
		start := time.Now()

		hash, nonce := ch.HashWorkingBlock()

		fmt.Printf("Added block 0x%x with nonce %d\n", hash, nonce)
		elapsed := time.Now().Sub(start).Milliseconds()
		fmt.Printf("%d hashes in %d ms\n", nonce, elapsed)
		rate := float64(nonce/uint64(elapsed)) / 1000.0
		fmt.Printf("%.3f Mh/s\n\n", rate)

		ch.AddWorkingBlock()
		ch.ResetWorkingBlock(hash)
	}
}

func (ch *Chain) Validate() bool {
	prev := make([]byte, block.HashLengthBytes)

	for i := 0; i < len(ch.blocks); i++ {
		if !bytes.Equal(ch.blocks[i].GetPrev(), prev) {
			return false
		}
		copy(prev, ch.blocks[i].GetHash())
	}

	// TODO: validate transactions

	return true
}
