# 🔗 GoChain

A complete blockchain implementation in Go featuring UTXO-based transactions, Proof of Work mining, P2P networking, and digital wallets.

[![Go Version](https://img.shields.io/badge/Go-1.13+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![GitHub](https://img.shields.io/badge/GitHub-parasraut21%2FGoChain-black.svg)](https://github.com/parasraut21/GoChain)

## 🌟 Features

- **🔐 Digital Wallets** - ECDSA-based wallet creation and management
- **💰 UTXO Model** - Unspent Transaction Output system for secure transactions
- **⛏️ Proof of Work** - Configurable mining difficulty with SHA-256 hashing
- **🌐 P2P Networking** - Decentralized node communication and synchronization
- **📊 Merkle Trees** - Efficient transaction verification and integrity
- **💾 Persistent Storage** - BadgerDB for reliable data persistence
- **🔒 Digital Signatures** - Cryptographic transaction signing and verification
- **📱 CLI Interface** - Easy-to-use command-line interface

## 🏗️ Architecture

```
GoChain/
├── main.go              # Application entry point
├── cli/                 # Command Line Interface
├── blockchain/          # Core blockchain logic
│   ├── block.go         # Block structure and creation
│   ├── blockchain.go    # Blockchain management
│   ├── transaction.go   # Transaction handling
│   ├── tx.go           # Transaction I/O structures
│   ├── utxo.go         # UTXO management
│   ├── merkle.go       # Merkle tree implementation
│   └── proof.go        # Proof of Work algorithm
├── wallet/              # Wallet and cryptography
│   ├── wallet.go       # Wallet structure and operations
│   ├── wallets.go      # Wallet collection management
│   └── utils.go        # Cryptographic utilities
├── network/             # P2P networking
│   └── network.go      # Network protocol and communication
└── tmp/                 # Database storage
```

## 🚀 Quick Start

### Prerequisites

- **Go 1.13+** - [Download Go](https://golang.org/dl/)
- **Git** - [Download Git](https://git-scm.com/downloads)

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/parasraut21/GoChain.git
   cd GoChain
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Build the project**
   ```bash
   go build
   ```

## 📖 Usage Guide

### 1. Set Environment Variable

**Windows (PowerShell):**
```powershell
$env:NODE_ID = "3000"
```

**Windows (CMD):**
```cmd
set NODE_ID=3000
```

**Linux/macOS:**
```bash
export NODE_ID=3000
```

### 2. Create Your First Wallet

```bash
go run main.go createwallet
```

**Expected Output:**
```
New address is: 1AT49VfoewnW9iQExVgHM8ZfxrGmysV7qX
```

### 3. Create a Blockchain

```bash
go run main.go createblockchain -address YOUR_ADDRESS
```

**Example:**
```bash
go run main.go createblockchain -address 1AT49VfoewnW9iQExVgHM8ZfxrGmysV7qX
```

**Expected Output:**
```
Genesis created
Finished!
```

### 4. Check Your Balance

```bash
go run main.go getbalance -address YOUR_ADDRESS
```

**Example:**
```bash
go run main.go getbalance -address 1AT49VfoewnW9iQExVgHM8ZfxrGmysV7qX
```

**Expected Output:**
```
Balance of 1AT49VfoewnW9iQExVgHM8ZfxrGmysV7qX: 20
```

### 5. Create Another Wallet

```bash
go run main.go createwallet
```

**Expected Output:**
```
New address is: 1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2
```

### 6. Send Coins

```bash
go run main.go send -from SENDER_ADDRESS -to RECEIVER_ADDRESS -amount AMOUNT -mine
```

**Example:**
```bash
go run main.go send -from 1AT49VfoewnW9iQExVgHM8ZfxrGmysV7qX -to 1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2 -amount 5 -mine
```

**Expected Output:**
```
Success!
```

### 7. Verify Transaction

Check the balance of both addresses:

```bash
go run main.go getbalance -address 1AT49VfoewnW9iQExVgHM8ZfxrGmysV7qX
go run main.go getbalance -address 1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2
```

**Expected Output:**
```
Balance of 1AT49VfoewnW9iQExVgHM8ZfxrGmysV7qX: 15
Balance of 1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2: 5
```

## 📋 Available Commands

| Command | Description | Example |
|---------|-------------|---------|
| `createwallet` | Create a new wallet | `go run main.go createwallet` |
| `listaddresses` | List all wallet addresses | `go run main.go listaddresses` |
| `createblockchain` | Create new blockchain | `go run main.go createblockchain -address ADDRESS` |
| `getbalance` | Check wallet balance | `go run main.go getbalance -address ADDRESS` |
| `send` | Send coins between addresses | `go run main.go send -from FROM -to TO -amount AMOUNT -mine` |
| `printchain` | Display the blockchain | `go run main.go printchain` |
| `reindexutxo` | Rebuild UTXO set | `go run main.go reindexutxo` |
| `startnode` | Start P2P node | `go run main.go startnode -miner ADDRESS` |

## 🔧 Advanced Usage

### View the Blockchain

```bash
go run main.go printchain
```

**Expected Output:**
```
Hash: 0000000000000000000000000000000000000000000000000000000000000000
Prev. hash: 
PoW: true
--- Transaction 1234567890abcdef...:
     Input 0:
       TXID:     0000000000000000000000000000000000000000000000000000000000000000
       Out:       -1
       Signature: 
       PubKey:    4669727374205472616e73616374696f6e2066726f6d2047656e65736973
     Output 0:
       Value:  20
       Script: 1234567890abcdef...
```

### List All Addresses

```bash
go run main.go listaddresses
```

**Expected Output:**
```
1AT49VfoewnW9iQExVgHM8ZfxrGmysV7qX
1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2
```

### Start a Mining Node

```bash
go run main.go startnode -miner YOUR_ADDRESS
```

**Expected Output:**
```
Starting Node 3000
Mining is on. Address to receive rewards: 1AT49VfoewnW9iQExVgHM8ZfxrGmysV7qX
```

## 🌐 P2P Network

### Start Multiple Nodes

**Terminal 1 (Node 3000):**
```bash
$env:NODE_ID = "3000"
go run main.go startnode -miner 1AT49VfoewnW9iQExVgHM8ZfxrGmysV7qX
```

**Terminal 2 (Node 3001):**
```bash
$env:NODE_ID = "3001"
go run main.go startnode
```

**Expected Output (Node 3000):**
```
Starting Node 3000
Mining is on. Address to receive rewards: 1AT49VfoewnW9iQExVgHM8ZfxrGmysV7qX
```

**Expected Output (Node 3001):**
```
Starting Node 3001
```

### Network Commands

The P2P network supports the following protocol commands:
- `version` - Node version and blockchain height synchronization
- `addr` - Known node address sharing
- `getblocks` - Request block hashes
- `getdata` - Request specific blocks or transactions
- `block` - Send block data
- `tx` - Send transaction data
- `inv` - Inventory (list of available items)

## 🔐 Security Features

### Digital Signatures
- **ECDSA P-256** elliptic curve cryptography
- **SHA-256** hashing for transaction integrity
- **RIPEMD-160** for address generation

### Proof of Work
- **Configurable difficulty** (currently set to 12)
- **SHA-256** mining algorithm
- **Nonce-based** block validation

### UTXO Model
- **One-time spending** - Each output can only be spent once
- **Change handling** - Automatic change calculation
- **Double-spending prevention** - Built-in protection

## 📊 Technical Specifications

| Feature | Implementation |
|---------|----------------|
| **Consensus** | Proof of Work (SHA-256) |
| **Difficulty** | 12 (configurable) |
| **Block Time** | Variable (based on difficulty) |
| **Transaction Model** | UTXO |
| **Cryptography** | ECDSA P-256, SHA-256, RIPEMD-160 |
| **Storage** | BadgerDB |
| **Networking** | TCP P2P |
| **Address Format** | Base58Check |

## 🧪 Testing

### Run Tests
```bash
go test ./...
```

### Test Coverage
```bash
go test -cover ./...
```

## 📁 Project Structure

```
GoChain/
├── 📄 main.go                 # Application entry point
├── 📁 cli/                    # Command Line Interface
│   └── 📄 cli.go             # CLI command handlers
├── 📁 blockchain/             # Core blockchain implementation
│   ├── 📄 block.go           # Block structure and creation
│   ├── 📄 blockchain.go      # Blockchain management
│   ├── 📄 transaction.go     # Transaction processing
│   ├── 📄 tx.go             # Transaction I/O structures
│   ├── 📄 utxo.go           # UTXO management
│   ├── 📄 merkle.go         # Merkle tree implementation
│   ├── 📄 proof.go          # Proof of Work algorithm
│   └── 📄 chain_iter.go     # Blockchain iterator
├── 📁 wallet/                 # Wallet and cryptography
│   ├── 📄 wallet.go         # Wallet structure and operations
│   ├── 📄 wallets.go        # Wallet collection management
│   └── 📄 utils.go          # Cryptographic utilities
├── 📁 network/                # P2P networking
│   └── 📄 network.go        # Network protocol implementation
├── 📁 tmp/                    # Database storage
│   └── 📁 blocks/           # BadgerDB data files
├── 📄 go.mod                 # Go module definition
├── 📄 go.sum                 # Go module checksums
└── 📄 README.md              # This file
```

## 🤝 Contributing

1. **Fork the repository**
2. **Create a feature branch** (`git checkout -b feature/amazing-feature`)
3. **Commit your changes** (`git commit -m 'Add some amazing feature'`)
4. **Push to the branch** (`git push origin feature/amazing-feature`)
5. **Open a Pull Request**

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👨‍💻 Author

**Paras Raut**
- GitHub: [@parasraut21](https://github.com/parasraut21)
- Project: [GoChain](https://github.com/parasraut21/GoChain)

## 🙏 Acknowledgments

- Inspired by Bitcoin's architecture
- Built with Go's excellent standard library
- Uses BadgerDB for persistent storage
- Implements industry-standard cryptographic algorithms

## 📞 Support

If you have any questions or need help:

1. **Check the documentation** above
2. **Open an issue** on GitHub
3. **Review existing issues** for solutions

---

**⭐ Star this repository if you found it helpful!**