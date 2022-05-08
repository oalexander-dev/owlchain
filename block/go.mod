module owlchain.localhost/block

go 1.18

replace owlchain.localhost/transaction => ../transaction

replace owlchain.localhost/address => ../address

require owlchain.localhost/transaction v0.0.0-00010101000000-000000000000

require owlchain.localhost/address v0.0.0-00010101000000-000000000000 // indirect
