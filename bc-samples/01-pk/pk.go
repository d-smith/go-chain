package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func main() {

	// Cribbing from https://github.com/ethereum/go-ethereum/blob/master/crypto/secp256k1/secp256_test.go
	key, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	privkey := make([]byte, 32)
	blob := key.D.Bytes()
	copy(privkey[32-len(blob):], blob)

	fmt.Println("Private key:", hex.EncodeToString(privkey))

	pubkey := elliptic.Marshal(secp256k1.S256(), key.X, key.Y)
	fmt.Println("Public key:", hex.EncodeToString(pubkey))
}
