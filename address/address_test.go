package address

import (
	"bytes"
	"testing"
)

func TestNewAddress(t *testing.T) {
	addr := NewAddress([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0})
	if bytes.Compare(addr.data, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}) != 0 {
		t.Fatal("Data did not match expected val after initializing.")
	}
}

func TestAddress_GetData(t *testing.T) {
	addr := NewAddress([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	if bytes.Compare(addr.GetData(), []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}) != 0 {
		t.Fatal("Data did not match expected val when retrieved with GetData.")
	}
}
