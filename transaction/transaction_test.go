package transaction

import (
	"bytes"
	addr "owlchain.localhost/address"
	"testing"
)

func TestNewTransaction(t *testing.T) {
	to := addr.NewAddress(make([]byte, addr.LengthBytes))
	from := addr.NewAddress(make([]byte, addr.LengthBytes))
	amount := uint64(5)
	nonce := uint64(1)
	tx := NewTransaction(*to, *from, amount, nonce)

	if bytes.Compare(tx.to.GetData(), to.GetData()) != 0 {
		t.Fatal("To did not match expected val after initializing.")
	}

	if bytes.Compare(tx.from.GetData(), from.GetData()) != 0 {
		t.Fatal("From did not match expected val after initializing.")
	}

	if tx.amount != amount {
		t.Fatal("Amount did not match expected val after initializing.")
	}

	if tx.nonce != nonce {
		t.Fatal("Nonce did not match expected val after initializing.")
	}
}

func TestTransaction_Serialize(t *testing.T) {
	to := addr.NewAddress(make([]byte, addr.LengthBytes))
	from := addr.NewAddress(make([]byte, addr.LengthBytes))
	amount := uint64(5)
	nonce := uint64(1)
	tx := NewTransaction(*to, *from, amount, nonce)

	sr := make([]byte, LengthBytes)

	copy(sr, tx.Serialize())
	txds := Deserialize(sr)

	equal := bytes.Equal(txds.to.GetData(), tx.to.GetData()) &&
		bytes.Equal(txds.from.GetData(), tx.from.GetData()) &&
		txds.amount == tx.amount && txds.nonce == tx.nonce

	if !equal {
		t.Fatal("Result of deserializing does not match original.")
	}
}
