package merkletree

import (
	"bytes"
	"errors"
)

// MerkleTree represents the merkle tree data structure. It holds references to the root of the tree,
// its leaf nodes, the merkle root hash and the type of hash function it supports.
type MerkleTree struct {
	Root           *Node
	Leafs          []*Node
	MerkleRootHash []byte
	HashFunc       HashFunc
}

// NewTree creates a new MerkleTree using provided payloads and a type of hash function.
func NewTree(pp []Payload, hashFunc HashFunc) (*MerkleTree, error) {
	t := &MerkleTree{
		HashFunc: hashFunc,
	}
	root, leafs, err := constructTreeFromPayloads(pp, t)

	if err != nil {
		return nil, err
	}

	t.Root = root
	t.Leafs = leafs
	t.MerkleRootHash = root.Hash

	return t, nil
}

// RebuildTree rebuilds the tree reusing only its leaf node payloads.
func (m *MerkleTree) RebuildTree() error {
	var pp []Payload

	for _, n := range m.Leafs {
		pp = append(pp, n.Payload)
	}

	return m.RebuildTreeWith(pp)
}

// RebuildTreeWith replaces the payloads of the tree and does a complete rebuild. No new
// tree instance is constructed, because the same instance is re-used.
func (m *MerkleTree) RebuildTreeWith(pp []Payload) error {
	root, leafs, err := constructTreeFromPayloads(pp, m)
	if err != nil {
		return err
	}

	m.Root = root
	m.Leafs = leafs
	m.MerkleRootHash = root.Hash

	return nil
}

// VerifyTree verifies the entire tree by validating the hashes at each tree level and returns true if the
// resulting hash at the root of the tree matches the merkle root hash.
func (m *MerkleTree) VerifyTree() (bool, error) {
	calculatedMerkleRoot, err := m.Root.verifyNode()
	if err != nil {
		return false, err
	}

	return bytes.Equal(m.MerkleRootHash, calculatedMerkleRoot), nil
}

// VerifyPayload checks whether a given payload is part of the tree and the hashes are valid for that payload.
// Returns true if the expected merkle root is equal to the merkle root calculated from the merkle path
// for a given payload. Returns true if valid and false otherwise.
func (m *MerkleTree) VerifyPayload(payload Payload) (bool, error) {
	for _, l := range m.Leafs {
		ok, err := l.Payload.Equals(payload)
		if err != nil {
			return false, err
		}

		if ok {
			currentParent := l.Parent
			for currentParent != nil {
				rightBytes, err := currentParent.Right.CalculateNodeHash()
				if err != nil {
					return false, err
				}

				leftBytes, err := currentParent.Left.CalculateNodeHash()
				if err != nil {
					return false, err
				}

				hashBytes, err := m.HashFunc.Calculate(append(leftBytes, rightBytes...))
				if err != nil {
					return false, err
				}

				if !bytes.Equal(hashBytes, currentParent.Hash) {
					return false, nil
				}

				currentParent = currentParent.Parent
			}

			return true, nil
		}
	}

	return false, nil
}

// GetMerklePath traces all the tree nodes needed for payload verification.
func (m *MerkleTree) GetMerklePath(payload Payload) ([][]byte, []int64, error) {
	for _, current := range m.Leafs {
		ok, err := current.Payload.Equals(payload)
		if err != nil {
			return nil, nil, err
		}

		if ok {
			currentParent := current.Parent

			var (
				merklePath [][]byte
				index      []int64
			)

			for currentParent != nil {
				if bytes.Equal(currentParent.Left.Hash, current.Hash) {
					merklePath = append(merklePath, currentParent.Right.Hash)
					index = append(index, 1) // right leaf
				} else {
					merklePath = append(merklePath, currentParent.Left.Hash)
					index = append(index, 0) // left leaf
				}

				current = currentParent
				currentParent = currentParent.Parent
			}

			return merklePath, index, nil
		}
	}

	return nil, nil, nil
}

// constructTreeFromPayloads constructs all levels given list of payloads until it reaches
// the root of the tree. Returns the resulting root node and a list of the leaf nodes.
func constructTreeFromPayloads(pp []Payload, tree *MerkleTree) (*Node, []*Node, error) {
	if len(pp) == 0 {
		return nil, nil, errors.New("error: cannot construct tree with no payload")
	}

	var leafNodes []*Node

	for _, p := range pp {
		hash, err := p.CalculateHash()
		if err != nil {
			return nil, nil, err
		}

		leafNodes = append(leafNodes, &Node{
			Hash:    hash,
			Payload: p,
			isLeaf:  true,
			Tree:    tree,
		})
	}

	leafNodesAreOddNumber := len(leafNodes)%2 == 1

	if leafNodesAreOddNumber {
		lastLeafNode := leafNodes[len(leafNodes)-1]

		duplicateLeafNode := &Node{
			Hash:        lastLeafNode.Hash,
			Payload:     lastLeafNode.Payload,
			isLeaf:      true,
			isDuplicate: true,
			Tree:        tree,
		}
		leafNodes = append(leafNodes, duplicateLeafNode)
	}

	root, err := constructNonLeafTreeLevelsFromLeafNodes(leafNodes, tree)
	if err != nil {
		return nil, nil, err
	}

	return root, leafNodes, nil
}

// constructNonLeafTreeLevelsFromLeafNodes constructs the non leaf tree levels given list of leaf nodes until it
// reaches the root of the tree. Returns the resulting root node.
func constructNonLeafTreeLevelsFromLeafNodes(leafNodes []*Node, tree *MerkleTree) (*Node, error) {
	var nodes []*Node

	for i := 0; i < len(leafNodes); i += 2 {
		var (
			left  int = i
			right int = i + 1
		)

		if i+1 == len(leafNodes) {
			right = i
		}

		hashBytes, err := tree.HashFunc.Calculate(
			append(leafNodes[left].Hash, leafNodes[right].Hash...),
		)
		if err != nil {
			return nil, err
		}

		n := &Node{
			Left:  leafNodes[left],
			Right: leafNodes[right],
			Hash:  hashBytes,
			Tree:  tree,
		}

		nodes = append(nodes, n)
		leafNodes[left].Parent = n
		leafNodes[right].Parent = n

		if len(leafNodes) == 2 {
			return n, nil
		}
	}

	return constructNonLeafTreeLevelsFromLeafNodes(nodes, tree)
}
