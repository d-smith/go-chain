package main

import (
	"log"

	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	tokenAddress := common.HexToAddress("0xdb98a5bFBA239000213813B2615B8A96E950A79B")
	instance, err := NewBalance(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x9949f7e672a568bB3EBEB777D5e8D1c1107e96E5")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bal)

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(18)))

	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"

}
