package transaction

import (
	"encoding/binary"
	addr "owlchain.localhost/address"
)

const LengthBytes = (2 * addr.LengthBytes) + (2 * 8)

type Transaction struct {
	to     addr.Address
	from   addr.Address
	amount uint64
	nonce  uint64
}

func NewTransaction(to addr.Address, from addr.Address, amount uint64, nonce uint64) *Transaction {
	tx := Transaction{to: to, from: from, amount: amount, nonce: nonce}
	return &tx
}

func (tx *Transaction) Serialize() []byte {
	sr := make([]byte, LengthBytes)

	copy(sr, tx.to.GetData())
	copy(sr[addr.LengthBytes:], tx.from.GetData())
	binary.PutUvarint(sr[(addr.LengthBytes*2):], tx.amount)
	binary.PutUvarint(sr[(addr.LengthBytes*2)+8:], tx.nonce)

	return sr
}

func Deserialize(sr []byte) Transaction {
	to := addr.NewAddress(sr[0:addr.LengthBytes])
	from := addr.NewAddress(sr[addr.LengthBytes:(addr.LengthBytes * 2)])
	amount, _ := binary.Uvarint(sr[(addr.LengthBytes * 2) : (addr.LengthBytes*2)+8])
	nonce, _ := binary.Uvarint(sr[(addr.LengthBytes*2)+8 : (addr.LengthBytes*2)+8+8])

	tx := NewTransaction(*to, *from, amount, nonce)
	return *tx
}
