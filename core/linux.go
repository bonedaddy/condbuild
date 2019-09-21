// +build linux

package core

import (
	"fmt"
	"math/big"
	"log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

// Linux does stuff
func Linux() {
	fmt.Println("linux")
	_, _ = NewBlockchain()
}

// NewBlockchain is used to generate a simulated blockchain
// this should not work for darwin due to gosigar deps.
func NewBlockchain() (*bind.TransactOpts, *backends.SimulatedBackend) {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(key)
	// https://medium.com/coinmonks/unit-testing-solidity-contracts-on-ethereum-with-go-3cc924091281
	balance := new(big.Int).Mul(big.NewInt(999999999999999), big.NewInt(999999999999999))
	gAlloc := map[common.Address]core.GenesisAccount{
		auth.From: {Balance: balance},
	}
	sim := backends.NewSimulatedBackend(gAlloc, 8000000)
	return auth, sim
}
