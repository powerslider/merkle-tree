package merkletree_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	merkletree "github.com/powerslider/merkle-tree"
)

var inputs = []struct {
	testCaseName   string
	payloads       []merkletree.Payload
	expectedHash   string
	invalidPayload merkletree.Payload
}{
	{
		testCaseName: "8 tx",
		payloads: []merkletree.Payload{
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "mznsVLaixRTMsdX4LuniyrMsuKqFh86yCW",
				ReceiverAddress: "mfem2dBeAnnnLZdjyKa33mm5zzVBpUShRB",
				Amount:          3.576564645,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "mytPq1KMUWvvBwMWsVTZUvHWaLqCwa2sdQ",
				ReceiverAddress: "mtkepLB437zfEHDx8m1YqKgtthVpkLV69d",
				Amount:          6.4675645,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "mjRNf5byc94Eh6URwV4MwpyVXcWKMq9fLp",
				ReceiverAddress: "ms2mR5yUTgAtSt7YLJT4StpkTYvSR1NB1J",
				Amount:          0.348,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "mvv25iaUDZE7JSbPTPW4jVjs7EK8mSkKHV",
				ReceiverAddress: "n1PJCAtwaHgeeJypATLYX7ck7jB9R5CpLG",
				Amount:          1.437,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "mkpH5XdvjoDJS5RN89obdkLhk2uLWFfB6e",
				ReceiverAddress: "n2zRDstggj6DHDDwPgB6dECBJ14U7sjgAi",
				Amount:          5.42342,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "muYrJuiNmy6M7CNT2McNcgMvRN1bhLCo2z",
				ReceiverAddress: "mt1Teo7ViAq1tMwfBq6Mnb4UhRMGQvZq62",
				Amount:          544.4456347,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "n2m7bbFC2HcPtRpwBJdQipsCVPRd19Z1XG",
				ReceiverAddress: "mptXBN66Qjswt4HnhfRdSfRG1dF87Uz7fp",
				Amount:          445.42342,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "minEKZwKaieGPk57p2BRJjs7DkqFJ15U7w",
				ReceiverAddress: "mftgohbV53HiqFNPmhez6ye1ovxJ8qrpMu",
				Amount:          245.423,
			},
		},
		expectedHash: "7659764119c1e53ff0247692c838fe3f0f8329f57db2727bd435602dda7155cf",
		invalidPayload: merkletree.PaymentTransactionPayload{
			SenderAddress:   "n1DTmXp5HqwH1Qrwd4fP1BB4zpBGDsxu2V",
			ReceiverAddress: "n4ZQsjdwNewrXXAy5C98VrYwT5MAAQ7nus",
			Amount:          123.2303,
		},
	},
	{
		testCaseName: "7 tx",
		payloads: []merkletree.Payload{
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "mqTNR1FK36xqwvE6UYoPwZHg6a6Mwdciav",
				ReceiverAddress: "mfem2dBeAnnnLZdjyKa33mm5zzVBpUShRB",
				Amount:          7.333335,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "mytPq1KMUWvvBwMWsVTZUvHWaLqCwa2sdQ",
				ReceiverAddress: "mv9a2b3biQbaZ4nPPKvuEpF2ut1rXRswBR",
				Amount:          6.053,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "mzxMjxnht5ZcCGkQwjxcM2YrjyS4wFaptu",
				ReceiverAddress: "ms2mR5yUTgAtSt7YLJT4StpkTYvSR1NB1J",
				Amount:          65.348,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "mvv25iaUDZE7JSbPTPW4jVjs7EK8mSkKHV",
				ReceiverAddress: "mnBtkLxzw9SKFLi8SKX16165wFiZ8Y8UYa",
				Amount:          3.437,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "mfevP1TpcTZr6ZDdGzFQvVSBmdLteyzoUC",
				ReceiverAddress: "miPgWajFJ54KYxTWTh9vFzRRRE6sfmMDDi",
				Amount:          15.423423890,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "mr5WZMkubVjiiU2qV22CyRZ8rWrNrmF6th",
				ReceiverAddress: "mt1Teo7ViAq1tMwfBq6Mnb4UhRMGQvZq62",
				Amount:          17.4347,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "mfevP1TpcTZr6ZDdGzFQvVSBmdLteyzoUC",
				ReceiverAddress: "mptXBN66Qjswt4HnhfRdSfRG1dF87Uz7fp",
				Amount:          88.42342,
			},
		},
		expectedHash: "37bbbcd13f425a9a3d6cce40570edc2baf2764e46bff15191874cb1a5035abc6",
		invalidPayload: merkletree.PaymentTransactionPayload{
			SenderAddress:   "mz43bGoaLXwcsen97SE8LLPUGF8wu3V5g1",
			ReceiverAddress: "mrvjMuJDts19vEqMNEJfEny3zfpy6GXY3Z",
			Amount:          101.23048933,
		},
	},
	{
		testCaseName: "5 tx",
		payloads: []merkletree.Payload{
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "n3vSXc2AiUXSzofqgkajTZuvG4DzHRmUtZ",
				ReceiverAddress: "mx88xwwNSJ7dgP6WucpnHLcKwnDvVxLBvB",
				Amount:          0.305,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "mi6VZfQFgeBaVeT4SZhfZcV3qpUdJ5tVTv",
				ReceiverAddress: "muECmHpKdDDLXiQZcdfRdC6Hc8m9A8pf3N",
				Amount:          1.0534905,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "mpYMT9z7Gj1yHP2fqtjevNVtgtMaYUmFqg",
				ReceiverAddress: "mrqW7MuNNY6t4V2EUA9TJMXJY5Uy3582hG",
				Amount:          54.345903,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "mjF5xMu2gxg35aWXVAa8nyjXxeBsdker73",
				ReceiverAddress: "mt1Teo7ViAq1tMwfBq6Mnb4UhRMGQvZq62",
				Amount:          10.4347,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "mfevP1TpcTZr6ZDdGzFQvVSBmdLteyzoUC",
				ReceiverAddress: "miPgWajFJ54KYxTWTh9vFzRRRE6sfmMDDi",
				Amount:          15.423423890,
			},
		},
		expectedHash: "9075827b24acac27b91839fd74f73cd1e50e64db9f5ea6681aa061adde420419",
		invalidPayload: merkletree.PaymentTransactionPayload{
			SenderAddress:   "myZTJihxZrQ1NxFSvJXCxnrgKTx8zYocQz",
			ReceiverAddress: "mmYG5VKVq4fBWekDoRVgJiPLbCiaxgNLt5",
			Amount:          23.2304893,
		},
	},
	{
		testCaseName: "4 tx",
		payloads: []merkletree.Payload{
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "n3vSXc2AiUXSzofqgkajTZuvG4DzHRmUtZ",
				ReceiverAddress: "mx88xwwNSJ7dgP6WucpnHLcKwnDvVxLBvB",
				Amount:          0.1234,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "mi6VZfQFgeBaVeT4SZhfZcV3qpUdJ5tVTv",
				ReceiverAddress: "muECmHpKdDDLXiQZcdfRdC6Hc8m9A8pf3N",
				Amount:          1.458495,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "moa4A5gLqSs1VsRpkiddXFuPTXkrSJ7ELP",
				ReceiverAddress: "mrqW7MuNNY6t4V2EUA9TJMXJY5Uy3582hG",
				Amount:          4.49503490,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "mjF5xMu2gxg35aWXVAa8nyjXxeBsdker73",
				ReceiverAddress: "mvEZWtsX6L9xWtWUsZQSxTtoAjkdhamX3D",
				Amount:          10.4347,
			},
		},
		expectedHash: "dcda87fa5a4f0ab2f5090223dd7150740786f3087deda58bbd4d808bfce0ea29",
		invalidPayload: merkletree.PaymentTransactionPayload{
			SenderAddress:   "mj9caWERxeG75bj5u1sAg8V1fRmeEGSU1z",
			ReceiverAddress: "mymVUoXH7mBvFGbnikKkQ1jTU2getk3Cwg",
			Amount:          1.5635435,
		},
	},
	{
		testCaseName: "3 tx",
		payloads: []merkletree.Payload{
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "n3NzHnLsfKvJVL3xEkAjuKBthpjteBZmaB",
				ReceiverAddress: "mjzDiq3fDhnqiGWV8vGvXaqi8z32zBHvvV",
				Amount:          4.138945,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "muSUQivVvLtR88WZgrrZkg1HiZHS7475tv",
				ReceiverAddress: "mqaQrxhzaWy4qmJMbN8qFHo4Vd3hXqrWXd",
				Amount:          6.443595,
			},
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "n1jAt5YDf7mhdAgh6iW6tTUB2AU6TQ71in",
				ReceiverAddress: "mvEZWtsX6L9xWtWUsZQSxTtoAjkdhamX3D",
				Amount:          15.3485390,
			},
		},
		expectedHash: "b1aa8a0967e4f12ba8aac0b0779e17bfb554b30f1322c3b09f5ed9eec5c5cc07",
		invalidPayload: merkletree.PaymentTransactionPayload{
			SenderAddress:   "mrfjThNkQing9L5TzhwjsiG2qyY9d6PmPr",
			ReceiverAddress: "mjKjgp4XLUKyBN87UiUKEQWxPgnP8ycJ7D",
			Amount:          39.34905,
		},
	},
	{
		testCaseName: "1 tx",
		payloads: []merkletree.Payload{
			merkletree.PaymentTransactionPayload{
				SenderAddress:   "n3vSXc2AiUXSzofqgkajTZuvG4DzHRmUtZ",
				ReceiverAddress: "mx88xwwNSJ7dgP6WucpnHLcKwnDvVxLBvB",
				Amount:          0.1234,
			},
		},
		expectedHash: "c16314270f5768f4845210c45e5797d3c53673952b5a780ccbc9bd31f48375cb",
		invalidPayload: merkletree.PaymentTransactionPayload{
			SenderAddress:   "mxZJmoDfbJYYfvGAM1WKKzjowkCTZetxqZ",
			ReceiverAddress: "mnh3nmg5dZ9YT474CaUKixovzK6K3YRrRD",
			Amount:          3.53453,
		},
	},
}

func TestNewTree(t *testing.T) {
	for _, test := range inputs {
		tree, err := merkletree.NewTree(test.payloads, merkletree.SHA256())
		if err != nil {
			t.Errorf("[test case: %s] error: unexpected error: %v", test.testCaseName, err)
		}

		expectedHashBytes, err := hex.DecodeString(test.expectedHash)
		if err != nil {
			t.Fatal("error hex decoding expected merkle root hash: ", err)
		}

		actualMerkleRootHash := hex.EncodeToString(tree.MerkleRootHash)

		if !bytes.Equal(tree.MerkleRootHash, expectedHashBytes) {
			t.Errorf("[test case: %s] error: expected hash equal to %v got %v",
				test.testCaseName, test.expectedHash, actualMerkleRootHash,
			)
		}
	}
}

func TestMerkleTreeVerifyTree(t *testing.T) {
	for _, test := range inputs {
		tree, err := merkletree.NewTree(test.payloads, merkletree.SHA256())
		if err != nil {
			t.Errorf("[test case: %s] error: unexpected error: %v", test.testCaseName, err)
		}

		isMerkleRootValid, err := tree.VerifyTree()
		if err != nil {
			t.Fatal("error verifying the merkle root hash: ", err)
		}

		if !isMerkleRootValid {
			t.Errorf("[test case: %s] error: expected tree to be valid", test.testCaseName)
		}

		tree.Root.Hash = []byte{123}
		tree.MerkleRootHash = []byte{123}

		isCompromisedMerkleRootValid, err := tree.VerifyTree()
		if err != nil {
			t.Fatal(err)
		}

		if isCompromisedMerkleRootValid {
			t.Errorf("[test case: %s] error: expected tree to be invalid", test.testCaseName)
		}
	}
}

func TestMerkleTreeVerifyPayload(t *testing.T) {
	for _, test := range inputs {
		tree, err := merkletree.NewTree(test.payloads, merkletree.SHA256())
		if err != nil {
			t.Errorf("[test case: %s] error: unexpected error: %v", test.testCaseName, err)
		}

		if len(test.payloads) > 0 {
			verifyValidPayload(t, tree, test.testCaseName, test.payloads[0])
		}

		if len(test.payloads) > 1 {
			verifyValidPayload(t, tree, test.testCaseName, test.payloads[1])
		}

		if len(test.payloads) > 2 {
			verifyValidPayload(t, tree, test.testCaseName, test.payloads[2])
		}

		if len(test.payloads) > 0 {
			verifyValidPayloadWithCompromisedMerkleRoot(t, tree, test.testCaseName, test.payloads[0])
		}

		verifyInvalidPayload(t, tree, test.testCaseName, test.invalidPayload)
	}
}

func TestMerkleTreeGetMerklePath(t *testing.T) {
	for _, test := range inputs {
		tree, err := merkletree.NewTree(test.payloads, merkletree.SHA256())
		if err != nil {
			t.Errorf("[test case: %s] error: unexpected error: %v", test.testCaseName, err)
		}

		for i := 0; i < len(test.payloads); i++ {
			merklePath, index, err := tree.GetMerklePath(test.payloads[i])
			if err != nil {
				t.Fatal(err)
			}

			merklePathHashBytes, err := tree.Leafs[i].CalculateNodeHash()
			if err != nil {
				t.Errorf("[test case: %s] error: calculateNodeHash error: %v", test.testCaseName, err)
			}

			for j := 0; j < len(merklePath); j++ {
				if index[j] == 1 {
					merklePathHashBytes = append(merklePathHashBytes, merklePath[j]...)
				} else {
					merklePathHashBytes = append(merklePath[j], merklePathHashBytes...)
				}

				merklePathHashBytes, err = tree.HashFunc.Calculate(merklePathHashBytes)
				if err != nil {
					t.Errorf("[test case: %s] error: Write error: %v", test.testCaseName, err)
				}
			}

			if !bytes.Equal(tree.MerkleRootHash, merklePathHashBytes) {
				t.Errorf(
					"[test case: %s] error: expected hash equal to %v got %v",
					test.testCaseName, merklePathHashBytes, tree.MerkleRootHash)
			}
		}
	}
}

func verifyValidPayload(t *testing.T, tree *merkletree.MerkleTree, testCaseName string, payload merkletree.Payload) {
	isPayloadValid, err := tree.VerifyPayload(payload)
	if err != nil {
		t.Fatal(err)
	}

	if !isPayloadValid {
		t.Errorf("[test case: %s] error: expected valid content", testCaseName)
	}
}

func verifyValidPayloadWithCompromisedMerkleRoot(
	t *testing.T, tree *merkletree.MerkleTree, testCaseName string, payload merkletree.Payload) {
	tree.Root.Hash = []byte{123}
	tree.MerkleRootHash = []byte{123}

	isCompromisedValid, err := tree.VerifyPayload(payload)
	if err != nil {
		t.Fatal(err)
	}

	if isCompromisedValid {
		t.Errorf("[test case: %s] error: expected invalid content", testCaseName)
	}

	if err := tree.RebuildTree(); err != nil {
		t.Fatal(err)
	}
}

func verifyInvalidPayload(
	t *testing.T, tree *merkletree.MerkleTree, testCaseName string, payload merkletree.Payload) {
	isInvalidPayloadValid, err := tree.VerifyPayload(payload)
	if err != nil {
		t.Fatal(err)
	}

	if isInvalidPayloadValid {
		t.Errorf("[test case: %s] error: expected invalid content", testCaseName)
	}
}
