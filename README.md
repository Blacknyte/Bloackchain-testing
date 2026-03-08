Simple Blockchain Implementation in Go
Student Information

Name: Muhammad Huzaifa Fayyaz
Roll Number: 22i-1692
Course: Blockchain and Cryptography Applications (BCA)
Assignment: Assignment 01 – Simple Blockchain Implementation

Assignment Objective

The objective of this assignment is to implement a simplified blockchain in Go to understand the fundamental concepts of blockchain technology. These concepts include block structure, hash linking, immutability, tamper detection, and blockchain verification logic.

This implementation focuses on the structural and verification aspects of blockchain and does not include a Proof-of-Work mechanism.

Blockchain Concepts Implemented

This project demonstrates the following blockchain principles:

Block Structure
Each block stores the following information:

Index

Timestamp

Transaction data

Nonce

Previous block hash

Current block hash

Hash Linking
Each block contains the hash of the previous block, which links all blocks together to form a chain.

Immutability
If the data of a block is modified, its hash changes. Since the next block stores the previous hash, this change breaks the chain integrity.

Tamper Detection
The blockchain verification process detects if any block has been altered.

Blockchain Verification
The VerifyChain() function recalculates block hashes and ensures the previous hash values match correctly.

Custom Requirements

Roll Number Injection
The last three digits of the roll number (692) are included inside one of the transaction strings.

Personalized Genesis Block
The Genesis Block is created with:

Transaction:
Genesis Block - 22i-1692

Nonce:
Sum of digits of roll number
1 + 6 + 9 + 2 = 18

Project Structure
assignment01bca_1692

│
├── main.go
│
└── assignment01bca
    └── blockchain.go

main.go
Responsible for creating blocks, printing the blockchain, verifying the chain, and demonstrating tampering.

blockchain.go
Contains the blockchain implementation including block structure and functions.

Implemented Functions

NewBlock()
Creates a new block and adds it to the blockchain.

ListBlocks()
Prints all blocks in the blockchain with their details.

ChangeBlock()
Modifies the transaction of a specific block to simulate tampering.

VerifyChain()
Checks whether the blockchain is valid by verifying hashes and previous hash links.

CalculateHash()
Generates a SHA-256 hash for a block using:

transaction + nonce + previousHash + index + timestamp

How to Run the Program

Step 1: Install Go
Download Go from:
https://go.dev/dl/

Step 2: Open the project folder in terminal.

Step 3: Run the program:

go run main.go
Program Demonstration

The program performs the following steps:

Creates the Genesis Block.

Adds two additional blocks to the blockchain.

Prints the complete blockchain.

Verifies that the blockchain is valid.

Modifies the transaction of Block 1.

Runs verification again to demonstrate tamper detection.

Expected behavior:

Before tampering
Blockchain is VALID

After tampering
Blockchain becomes INVALID due to hash mismatch.

Hash Algorithm Used

SHA-256 cryptographic hash function is used to generate block hashes.

This ensures:

Deterministic hashing

Fixed-length hash output

Tamper detection capability

Conclusion

This assignment demonstrates how blockchain ensures data integrity through cryptographic hashing and hash linking. By modifying a block and verifying the chain again, the system successfully detects tampering, illustrating the immutability property of blockchain systems.
