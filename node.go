package merkletree

import "fmt"

// Node represents a node, root, or leaf in the tree. It stores pointers to its immediate
// relationships, a hash, the content stored if it is a leaf, and other metadata.
type Node struct {
	Tree        *MerkleTree
	Parent      *Node
	Left        *Node
	Right       *Node
	isLeaf      bool
	isDuplicate bool
	Hash        []byte
	Payload     Payload
}

// String returns a string representation of the node.
func (n *Node) String() string {
	return fmt.Sprintf("%t %t %v %s", n.isLeaf, n.isDuplicate, n.Hash, n.Payload)
}

// verifyNode walks down the tree until hitting a leaf, calculating the hash at each level
// and returning the resulting hash of Node n.
func (n *Node) verifyNode() ([]byte, error) {
	if n.isLeaf {
		return n.Payload.CalculateHash()
	}

	rightBytes, err := n.Right.verifyNode()
	if err != nil {
		return nil, err
	}

	leftBytes, err := n.Left.verifyNode()
	if err != nil {
		return nil, err
	}

	return n.Tree.HashFunc.Calculate(append(leftBytes, rightBytes...))
}

// CalculateNodeHash is a helper function that calculates the hash of the node.
func (n *Node) CalculateNodeHash() ([]byte, error) {
	if n.isLeaf {
		return n.Payload.CalculateHash()
	}

	return n.Tree.HashFunc.Calculate(append(n.Left.Hash, n.Right.Hash...))
}
