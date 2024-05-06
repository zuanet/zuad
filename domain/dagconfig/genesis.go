// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
package dagconfig

import (
	"github.com/kaspanet/go-muhash"
	"github.com/zuanet/zuad/domain/consensus/model/externalapi"
	"github.com/zuanet/zuad/domain/consensus/utils/blockheader"
	"github.com/zuanet/zuad/domain/consensus/utils/subnetworks"
	"github.com/zuanet/zuad/domain/consensus/utils/transactionhelper"
	"math/big"
)

var genesisTxOuts = []*externalapi.DomainTransactionOutput{}

var genesisTxPayload = []byte{
	0xCE, 0x8D, 0x0F, 0x0D, 0x7E, 0x5D, 0x3C, 0xAD, 
	0x97, 0xBD, 0x8B, 0x3D, 0x24, 0x2C, 0x4C, 0xA2, 
	0x99, 0xDB, 0x6C, 0x89, 0x62, 0x6A, 0xA0, 0xCA, 
	0x72, 0x4B, 0xE4, 0x66, 0x4A, 0xEB, 0x43, 0x96,
}

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network.
var genesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0, []*externalapi.DomainTransactionInput{}, genesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, genesisTxPayload)

// genesisHash is the hash of the first block in the block DAG for the main
// network (genesis block).
var genesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x26, 0x80, 0xCE, 0x07, 0xAC, 0x8F, 0xED, 0x07,
	0x0C, 0x81, 0xD7, 0x0F, 0xF9, 0xCD, 0x6B, 0x77, 
	0xB4, 0x49, 0xC6, 0xC3, 0x0D, 0x06, 0x24, 0x0F, 
	0x57, 0x8E, 0x2C, 0xED, 0x4B, 0xA3, 0x6E, 0xD7,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xCE, 0x8D, 0x0F, 0x0D, 0x7E, 0x5D, 0x3C, 0xAD, 
	0x97, 0xBD, 0x8B, 0x3D, 0x24, 0x2C, 0x4C, 0xA2, 
	0x99, 0xDB, 0x6C, 0x89, 0x62, 0x6A, 0xA0, 0xCA, 
	0x72, 0x4B, 0xE4, 0x66, 0x4A, 0xEB, 0x43, 0x96,
})

// genesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the main network.
var genesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		genesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
			0x14d2bd72185,
        	525264379,
        	0x1a14e,
        	0, // Checkpoint DAA score
        	0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{genesisCoinbaseTx},
}

var devnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var devnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, //script version
	0x01,                                           // Varint
	0x00,                                           // OP-FALSE
	0x73, 0x65, 0x64, 0x72, 0x61, 0x2D, 0x6D, 0x61, 0x69, 0x6E, 0x6E, 0x65, 0x74, // sedra mainnet payload
}

// devnetGenesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the development network.
var devnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, devnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, devnetGenesisTxPayload)

// devGenesisHash is the hash of the first block in the block DAG for the development
// network (genesis block).
var devnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x7c, 0x38, 0x0a, 0xe6, 0x90, 0xf9, 0xea, 0xa6, 0x4f, 0x30, 0x04, 0x51, 0xd5, 0xe4, 0x66, 0x37, 0x11, 0x13, 0xc3, 0x8b, 0x96, 0xed, 0x96, 0xfe, 0x3f, 0x1e, 0x93, 0x48, 0x69, 0x35, 0x5f, 0xfa,
})

// devnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var devnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x0a, 0x65, 0x43, 0x55, 0xc1, 0xd5, 0x8b, 0xd4, 0xab, 0xd9, 0xbd, 0xcc, 0x2e, 0x52, 0xf1, 0xbd, 0x56, 0x58, 0x64, 0xe8, 0xa1, 0x6c, 0xe0, 0x37, 0x7b, 0x48, 0xec, 0xc6, 0x4d, 0xb5, 0x4b, 0x27,
})

// devnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var devnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		devnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
			0x14d2bd72185,
        	525264379,
        	0x1a14e,
        	0, // Checkpoint DAA score
        	0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{devnetGenesisCoinbaseTx},
}

var simnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var simnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                                   // Varint
	0x00,                                                                   // OP-FALSE
	0x6b, 0x61, 0x73, 0x70, 0x61, 0x2d, 0x73, 0x69, 0x6d, 0x6e, 0x65, 0x74, // sedra-simnet
}

// simnetGenesisCoinbaseTx is the coinbase transaction for the simnet genesis block.
var simnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, simnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, simnetGenesisTxPayload)

// simnetGenesisHash is the hash of the first block in the block DAG for
// the simnet (genesis block).
var simnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x41, 0x1f, 0x8c, 0xd2, 0x6f, 0x3d, 0x41, 0xae,
	0xa3, 0x9e, 0x78, 0x57, 0x39, 0x27, 0xda, 0x24,
	0xd2, 0x39, 0x95, 0x70, 0x5b, 0x57, 0x9f, 0x30,
	0x95, 0x9b, 0x91, 0x27, 0xe9, 0x6b, 0x79, 0xe3,
})

// simnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var simnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x19, 0x46, 0xd6, 0x29, 0xf7, 0xe9, 0x22, 0xa7,
	0xbc, 0xed, 0x59, 0x19, 0x05, 0x21, 0xc3, 0x77,
	0x1f, 0x73, 0xd3, 0x52, 0xdd, 0xbb, 0xb6, 0x86,
	0x56, 0x4a, 0xd7, 0xfd, 0x56, 0x85, 0x7c, 0x1b,
})

// simnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var simnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		simnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x17c5f62fbb6,
		0x207fffff,
		0x2,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{simnetGenesisCoinbaseTx},
}

var testnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var testnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                                         // Varint
	0x00,                                                                         // OP-FALSE
	0x6b, 0x61, 0x73, 0x70, 0x61, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x6e, 0x65, 0x74, // sedra-testnet
}

// testnetGenesisCoinbaseTx is the coinbase transaction for the testnet genesis block.
var testnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, testnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, testnetGenesisTxPayload)

// testnetGenesisHash is the hash of the first block in the block DAG for the test
// network (genesis block).
var testnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xf8, 0x96, 0xa3, 0x03, 0x48, 0x73, 0xbe, 0x17,
	0x39, 0xfc, 0x43, 0x59, 0x23, 0x68, 0x99, 0xfd,
	0x3d, 0x65, 0xd2, 0xbc, 0x94, 0xf9, 0x78, 0x0d,
	0xf0, 0xd0, 0xda, 0x3e, 0xb1, 0xcc, 0x43, 0x70,
})

// testnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for testnet.
var testnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x17, 0x34, 0x14, 0x08, 0xa5, 0x72, 0x45, 0x56,
	0x50, 0x4d, 0xf4, 0xd6, 0xcf, 0x51, 0x5c, 0xbf,
	0xbb, 0x22, 0x04, 0x30, 0xdc, 0x45, 0x1c, 0x74,
	0x3c, 0x22, 0xd5, 0xe9, 0x11, 0x72, 0x0c, 0x2a,
})

// testnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for testnet.
var testnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		testnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x17c5f62fbb6,
		0x1e7fffff,
		0x14582,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{testnetGenesisCoinbaseTx},
}

// devnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var devnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		devnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1714543200000,
               45552,
               0x41d00ff,
               1312860,
                0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{devnetGenesisCoinbaseTx},
}

var simnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var simnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                                   // Varint
	0x00,                                                                   // OP-FALSE
	0x6b, 0x61, 0x73, 0x70, 0x61, 0x2d, 0x73, 0x69, 0x6d, 0x6e, 0x65, 0x74, // zua-simnet
}

// simnetGenesisCoinbaseTx is the coinbase transaction for the simnet genesis block.
var simnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, simnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, simnetGenesisTxPayload)

// simnetGenesisHash is the hash of the first block in the block DAG for
// the simnet (genesis block).
var simnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x41, 0x1f, 0x8c, 0xd2, 0x6f, 0x3d, 0x41, 0xae,
	0xa3, 0x9e, 0x78, 0x57, 0x39, 0x27, 0xda, 0x24,
	0xd2, 0x39, 0x95, 0x70, 0x5b, 0x57, 0x9f, 0x30,
	0x95, 0x9b, 0x91, 0x27, 0xe9, 0x6b, 0x79, 0xe3,
})

// simnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var simnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x19, 0x46, 0xd6, 0x29, 0xf7, 0xe9, 0x22, 0xa7,
	0xbc, 0xed, 0x59, 0x19, 0x05, 0x21, 0xc3, 0x77,
	0x1f, 0x73, 0xd3, 0x52, 0xdd, 0xbb, 0xb6, 0x86,
	0x56, 0x4a, 0xd7, 0xfd, 0x56, 0x85, 0x7c, 0x1b,
})

// simnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var simnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		simnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1714543200000,
                45552,
                0x41d00ff,
                1312860,
                0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{simnetGenesisCoinbaseTx},
}

var testnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var testnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                                         // Varint
	0x00,                                                                         // OP-FALSE
	0x6b, 0x61, 0x73, 0x70, 0x61, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x6e, 0x65, 0x74, // zua-testnet
}

// testnetGenesisCoinbaseTx is the coinbase transaction for the testnet genesis block.
var testnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, testnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, testnetGenesisTxPayload)

// testnetGenesisHash is the hash of the first block in the block DAG for the test
// network (genesis block).
var testnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xf8, 0x96, 0xa3, 0x03, 0x48, 0x73, 0xbe, 0x17,
	0x39, 0xfc, 0x43, 0x59, 0x23, 0x68, 0x99, 0xfd,
	0x3d, 0x65, 0xd2, 0xbc, 0x94, 0xf9, 0x78, 0x0d,
	0xf0, 0xd0, 0xda, 0x3e, 0xb1, 0xcc, 0x43, 0x70,
})

// testnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for testnet.
var testnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x17, 0x34, 0x14, 0x08, 0xa5, 0x72, 0x45, 0x56,
	0x50, 0x4d, 0xf4, 0xd6, 0xcf, 0x51, 0x5c, 0xbf,
	0xbb, 0x22, 0x04, 0x30, 0xdc, 0x45, 0x1c, 0x74,
	0x3c, 0x22, 0xd5, 0xe9, 0x11, 0x72, 0x0c, 0x2a,
})

// testnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for testnet.
var testnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		testnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1714543200000,
                45552,
               0x41d00ff,
                1312860,
                0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{testnetGenesisCoinbaseTx},
}
