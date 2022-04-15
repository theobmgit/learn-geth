package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	for i := 14544170; i < 14544179; i++ {
		block, err := client.BlockByNumber(context.Background(), big.NewInt(14544179))
		if err != nil {
			log.Fatal(err)
		}

		log.Println(len(block.Transactions()))
	}
}
