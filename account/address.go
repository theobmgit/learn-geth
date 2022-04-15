package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
	"reflect"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")

	account := common.HexToAddress("0x00000000219ab540356cBB839Cbe05303d7705Fa")

	balance, err2 := client.BalanceAt(context.Background(), account, big.NewInt(14543005))
	if err2 != nil {
		log.Fatal(err2)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(ethValue)
	log.Println(reflect.TypeOf(balance))
	log.Println(reflect.TypeOf(ethValue))
}
