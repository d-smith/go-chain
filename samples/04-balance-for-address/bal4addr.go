package main

import (
    "context"
    "fmt"
    "log"
	"os"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	var (
		rpcEndpoint = os.Getenv("RPC_ENDPOINT")
		walletAddress0 = os.Getenv("WALLET_ADDRESS0")
	)

 	client, err := ethclient.Dial(rpcEndpoint)
    if err != nil {
        log.Fatal(err)
    }

	//HD wallet account 0 address
	account := common.HexToAddress(walletAddress0)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)
}