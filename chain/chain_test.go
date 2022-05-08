package chain

import "testing"

func TestChain_Run(t *testing.T) {
	ch := NewChain(2)
	ch.Run()
}
