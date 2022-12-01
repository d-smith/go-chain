package main

import (
    "fmt"
    "log"
    "os"

    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    var rpcEndpoint = os.Getenv("RPC_ENDPOINT")

	client, err := ethclient.Dial(rpcEndpoint)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("we have a connection")
    _ = client // we'll use this in the upcoming sections
}