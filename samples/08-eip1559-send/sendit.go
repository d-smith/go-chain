package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	var (
		rpcEndpoint    = os.Getenv("RPC_ENDPOINT")
		walletAccount0 = os.Getenv("WALLET_ACCOUNT0")
		account0PK     = os.Getenv("ACCOUNT0_PK")
		chainIDFromEnv = os.Getenv("CHAIN_ID")
	)

	client, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	// private key is for a prefunded ganache account
	privateKey, err := crypto.HexToECDSA(account0PK)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println(fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	tip := big.NewInt(2000000000)            // maxPriorityFeePerGas = 2 Gwei
	feeCap := big.NewInt(20000000000)        // maxFeePerGas = 20 Gwei
	if err != nil {
		log.Fatal(err)
	}

	// account 0 from hd wallet
	toAddress := common.HexToAddress(walletAccount0)
	var data []byte

	chainID := new(big.Int)
	chainID, ok = chainID.SetString(chainIDFromEnv, 10)
	if !ok {
		log.Fatal("error setting chain id from env")
	}
	fmt.Println(chainID)

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasFeeCap: feeCap,
		GasTipCap: tip,
		Gas:       gasLimit,
		To:        &toAddress,
		Value:     value,
		Data:      data,
	})

	fmt.Println("sign tx")
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("send tx")
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
