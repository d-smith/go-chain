package main


import (
	"bytes"
    "crypto/ecdsa"
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// get private key
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
	log.Fatal(err)
	}

	// hash some data
	data := []byte("hello")
    hash := crypto.Keccak256Hash(data)
    fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	// sign it
    signature, err := crypto.Sign(hash.Bytes(), privateKey)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(hexutil.Encode(signature)) 

	// public key
	publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	
	// verify it
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
    if err != nil {
        log.Fatal(err)
    }

    matches := bytes.Equal(sigPublicKey, publicKeyBytes)
    fmt.Println(matches) // true

    sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
    if err != nil {
        log.Fatal(err)
    }

    sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
    matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
    fmt.Println(matches) // true

    signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
    verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
    fmt.Println(verified) // true
}