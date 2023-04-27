# Merkle Tree

## Overview

A Merkle tree, also known as a hash tree, is a data structure that is used to efficiently verify the integrity and consistency of large datasets. It was invented by Ralph Merkle in 1979.

A Merkle tree consists of a tree of hash values, where each leaf node represents a small portion of the original dataset, such as a block of data or a transaction. The hash of each leaf node is calculated using a cryptographic hash function, such as SHA-256, which produces a fixed-length output that represents the data in a unique way.

The Merkle tree is constructed by repeatedly hashing pairs of adjacent leaf nodes, combining them into intermediate nodes, and then hashing those nodes together until a single root hash is produced. The root hash is a unique identifier for the entire dataset, and any change to any leaf node will result in a completely different root hash.

One important use of Merkle trees is in the context of blockchain technology. In a blockchain, each block contains a set of transactions, and the block header includes a Merkle root hash that represents the entire set of transactions in the block. This allows anyone to quickly and easily verify that the contents of a block have not been tampered with, without having to download and verify the entire set of transactions.

Another advantage of Merkle trees is that they allow for efficient partial verification of the dataset. By only downloading a portion of the tree and verifying the hash values, a user can determine whether a particular leaf node is part of the dataset or not, without having to download the entire dataset. This can be useful in situations where only a small portion of the dataset is needed, or where the entire dataset is too large to download in its entirety.

Overall, Merkle trees provide a robust and efficient way to verify the integrity and consistency of large datasets, making them a useful tool in a wide range of applications.

## Getting Started

**Step 1.** Install external tooling (golangci-lint, etc.):

```shell script
make install
```

**Step 2.** Run code linting and tests:

```shell script
make all
```

## Development Setup

**Step 0.** Install [pre-commit](https://pre-commit.com/):

```shell
pip install pre-commit

# For macOS users.
brew install pre-commit
```

Then run `pre-commit install` to setup git hook scripts.
Used hooks can be found [here](.pre-commit-config.yaml).

______________________________________________________________________

NOTE

> `pre-commit` aids in running checks (end of file fixing,
> markdown linting, go linting, runs go tests, json validation, etc.)
> before you perform your git commits.

______________________________________________________________________

**Step 1.** Install external tooling (golangci-lint, etc.):

```shell script
make install
```

**Step 2.** Run either of these command (depending on the use case):

```shell script
# Runs code linting and all tests.
make all

# Runs golangci-lint linter.
make lint

# Runs all tests.
make test

# Runs all tests with a code coverage report.
make cover-test
```

## License

[MIT](LICENSE)
