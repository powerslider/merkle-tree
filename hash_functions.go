package merkletree

import (
	"crypto/sha256"
	"hash"
)

// HashFunc represents a function returning the constructor of a specific type of hash algorithm.
type HashFunc func() hash.Hash

// Calculate calculates the hash of given payload using the specified hash algorithm.
func (h HashFunc) Calculate(payload []byte) ([]byte, error) {
	hFunc := h()
	if _, err := hFunc.Write(payload); err != nil {
		return nil, err
	}

	return hFunc.Sum(nil), nil
}

// SHA256 returns the constructor function for the SHA-256 algorithm.
func SHA256() HashFunc {
	return HashFunc(sha256.New)
}
