package block

import (
	"encoding/binary"
	trans "owlchain.localhost/transaction"
)

const HashLengthBytes = 32
const MaxTransactions = 64
const LengthBytes = (2 * HashLengthBytes) + 8 + (MaxTransactions * trans.LengthBytes)

type Block struct {
	hash            []byte
	prev            []byte
	nonce           uint64
	transactions    []trans.Transaction
	numTransactions uint16
}

func NewBlock() *Block {
	bk := Block{hash: make([]byte, HashLengthBytes), prev: make([]byte, HashLengthBytes),
		nonce: 0, transactions: make([]trans.Transaction, MaxTransactions), numTransactions: 0}

	return &bk
}

func NewBlockWithPrev(prev []byte) *Block {
	bk := Block{hash: make([]byte, HashLengthBytes), prev: make([]byte, HashLengthBytes),
		nonce: 0, transactions: make([]trans.Transaction, MaxTransactions), numTransactions: 0}

	copy(bk.prev, prev)

	return &bk
}

func (bk *Block) Serialize(includeHash bool) []byte {
	sr := make([]byte, LengthBytes)

	if includeHash {
		copy(sr, bk.hash)
	}

	copy(sr[HashLengthBytes:], bk.prev)
	binary.PutUvarint(sr[(HashLengthBytes*2):], bk.nonce)
	binary.BigEndian.PutUint16(sr[(HashLengthBytes*2)+8:], bk.numTransactions)

	offset := (HashLengthBytes * 2) + 8 + 16
	for i := uint16(0); i < bk.numTransactions; i++ {
		d := offset + (int(i) * trans.LengthBytes)
		tx := bk.transactions[i]
		copy(sr[d:], tx.Serialize())
	}

	return sr
}

func Deserialize(sr []byte) *Block {
	bk := NewBlock()

	copy(bk.hash, sr[0:HashLengthBytes])
	copy(bk.prev, sr[HashLengthBytes:(HashLengthBytes*2)])
	bk.nonce, _ = binary.Uvarint(sr[(HashLengthBytes * 2) : (HashLengthBytes*2)+8])
	bk.numTransactions = binary.BigEndian.Uint16(sr[(HashLengthBytes*2)+8 : (HashLengthBytes*2)+8+2])

	offset := (HashLengthBytes * 2) + 8 + 16
	for i := uint16(0); i < bk.numTransactions; i++ {
		d := offset + (int(i) * trans.LengthBytes)
		tx := trans.Deserialize(sr[d:(d + trans.LengthBytes)])
		bk.AddTransaction(tx)
	}

	return bk
}

func (bk *Block) UpdateHash(newHash []byte, nonce uint64) {
	copy(bk.hash, newHash)
	bk.nonce = nonce
}

func (bk *Block) AddTransaction(tx trans.Transaction) {
	bk.transactions[bk.numTransactions] = tx
	bk.numTransactions++
}

func (bk *Block) GetHash() []byte {
	return bk.hash
}

func (bk *Block) GetPrev() []byte {
	return bk.prev
}
