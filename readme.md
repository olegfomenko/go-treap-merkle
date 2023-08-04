# Treap-based dynamic Merkle tree

Proof of concept for dynamic Merkle tree. 
Treap as primary data structure used ao achieve logarithmic inserting/removing.

## Usage
```go
package main

import (
	"math/rand"
	
	merkle "github.com/olegfomenko/go-treap-merkle"
)

func main() {
	tree := merkle.New()
	
	// Insert
	tree.Insert([]byte("hash"), rand.Uint64())
	
	// Merkle path
	path := tree.MerklePath([]byte("hash"))
	
	// Merkle Root
	root := tree.MerkleRoot()

	// Remove
	tree.Remove([]byte("hash"))
}
```

## Related docs

Treap description: "<https://en.wikipedia.org/wiki/Treap>"

Merkle tree description: "<https://en.wikipedia.org/wiki/Merkle_tree>"

<img src="./concept.jpg" alt="Signing scheme" style="width:600px;"/>