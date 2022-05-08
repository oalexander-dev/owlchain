package address

const LengthBytes = 20

type Address struct {
	data []byte
}

func NewAddress(address []byte) *Address {
	a := Address{data: make([]byte, LengthBytes)}
	copy(a.data, address)
	return &a
}

func (addr *Address) GetData() []byte {
	return addr.data
}
