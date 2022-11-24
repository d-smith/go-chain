package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)
23
func main() {

 	client, err := ethclient.Dial("http://172.17.144.1:7545")
    if err != nil {
        log.Fatal(err)
    }

	//HD wallet account 0 address
	account := common.HexToAddress("0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)
}