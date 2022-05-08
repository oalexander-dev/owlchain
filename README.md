# Owlchain
A simple example of a Proof-of-Work blockchain.

## Entities
### Address
This package defines the address and its interface.

| Method                     | Description                                    |
|----------------------------|------------------------------------------------|
| NewAddress(address []byte) | Creates a new address using the provided value |
| (addr *Address) GetData() []byte                 | Returns the bytes of the address               |