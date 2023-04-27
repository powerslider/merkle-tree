package merkletree

import (
	"encoding/json"
	"reflect"
)

// Payload represents the data that is stored and verified by the tree. A type that
// implements this interface can be used as an item in the tree.
type Payload interface {
	CalculateHash() ([]byte, error)
	Equals(other Payload) (bool, error)
}

// PaymentTransactionPayload implements the Payload interface and represents the Payload stored in the tree.
// This implementation represents a payment transaction.
type PaymentTransactionPayload struct {
	SenderAddress   string  `json:"sender_address"`
	ReceiverAddress string  `json:"receiver_address"`
	Amount          float64 `json:"amount"`
}

// CalculateHash calculates the hash of the values of a PaymentTransactionPayload.
func (t PaymentTransactionPayload) CalculateHash() ([]byte, error) {
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return SHA256().Calculate(jsonBytes)
}

// Equals checks if two PaymentTransactionPayloads are equal.
func (t PaymentTransactionPayload) Equals(other Payload) (bool, error) {
	return reflect.DeepEqual(t, other), nil
}
