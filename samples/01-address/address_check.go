package main

import (
    "context"
    "fmt"
    "log"
    "regexp"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

    fmt.Printf("is valid: %v\n", re.MatchString("0x80902fbdb3DB2b97151c7E104dFee7B7aA43e51C")) // is valid: true
    fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false
	fmt.Printf("is valid: %v\n", re.MatchString("0xeEa6F93758Ebd196C5F2262893cCbDfE61d20626"))


    //client, err := ethclient.Dial("http://172.17.144.1:7545")
    client, err := ethclient.Dial("http://127.0.0.1:8545")
    if err != nil {
        log.Fatal(err)
    }

    // first ganache acct
    address := common.HexToAddress("0x80902fbdb3DB2b97151c7E104dFee7B7aA43e51C")
    bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
    if err != nil {
        log.Fatal(err)
    }

    isContract := len(bytecode) > 0

    fmt.Printf("is contract: %v\n", isContract) // is contract: fakse

    // 2nd user account address
    address = common.HexToAddress("0xeEa6F93758Ebd196C5F2262893cCbDfE61d20626")
    bytecode, err = client.CodeAt(context.Background(), address, nil) // nil is latest block
    if err != nil {
        log.Fatal(err)
    }

    isContract = len(bytecode) > 0

    fmt.Printf("is contract: %v\n", isContract) // is contract: false

	account := common.HexToAddress("0xeEa6F93758Ebd196C5F2262893cCbDfE61d20626")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)
}