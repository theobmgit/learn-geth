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
		log.Fatalln(err)
	}

	block, err := client.BlockByNumber(context.Background(), big.NewInt(14544369))
	if err != nil {
		log.Fatalln(err)
	}

	tx := block.Transactions()[1]
	receipt, _ := client.TransactionReceipt(context.Background(), tx.Hash())
	log.Println(receipt.Logs)
}
