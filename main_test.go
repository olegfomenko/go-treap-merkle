package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestTreap(t *testing.T) {
	Insert(crypto.Keccak256([]byte{1}), rand.Uint64())
	Insert(crypto.Keccak256([]byte{2}), rand.Uint64())
	Insert(crypto.Keccak256([]byte{3}), rand.Uint64())
	Insert(crypto.Keccak256([]byte{4}), rand.Uint64())
	Insert(crypto.Keccak256([]byte{5}), rand.Uint64())

	fmt.Println("Root: ", hexutil.Encode(Root.MerkleHash))

	path := MerklePath(crypto.Keccak256([]byte{3}))
	check := crypto.Keccak256([]byte{3})

	fmt.Println("Path: ")
	for _, p := range path {
		fmt.Println(hexutil.Encode(p))
	}

	for i := len(path) - 1; i >= 0; i-- {
		if len(path[i]) > 0 {
			check = hash(check, path[i])
		}
	}

	fmt.Println("Check:", hexutil.Encode(check), "should be equal to root:", hexutil.Encode(Root.MerkleHash))
}
