package block

import (
	"bytes"
	addr "owlchain.localhost/address"
	trans "owlchain.localhost/transaction"
	"testing"
)

func TestNewBlock(t *testing.T) {
	bk := NewBlock()
	if !bytes.Equal(bk.hash, make([]byte, HashLengthBytes)) ||
		!bytes.Equal(bk.prev, make([]byte, HashLengthBytes)) ||
		bk.numTransactions != 0 || bk.nonce != 0 {
		t.Fatal("New block had incorrect values after initializing.")
	}
}

func TestBlock_AddTransaction(t *testing.T) {
	to := addr.NewAddress(make([]byte, addr.LengthBytes))
	from := addr.NewAddress(make([]byte, addr.LengthBytes))
	amount := uint64(5)
	nonce := uint64(1)
	tx := trans.NewTransaction(*to, *from, amount, nonce)

	bk := NewBlock()
	bk.AddTransaction(*tx)

	if bk.numTransactions != 1 {
		t.Fatal("Num transactions did not increment.")
	}
}

func TestBlock_Serialize(t *testing.T) {
	bk := NewBlock()
	bk.nonce = 5

	sr := make([]byte, LengthBytes)
	copy(sr, bk.Serialize(true))

	bkds := Deserialize(sr)

	if bkds.nonce != bk.nonce {
		t.Fatal("Nonce did not match after deserializing block.")
	}
}