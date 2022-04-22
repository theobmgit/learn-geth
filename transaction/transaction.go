package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/ee09af9e74a243c4a7ee638fcf0a9d21")
	if err != nil {
		log.Fatalln(err)
	}

	rawTx := "0x3a497e8893df6201d6598a03116f22014b1e53dbd343ccf9048e1e77163dcc41"
	tx, pending, _ := client.TransactionByHash(context.Background(), common.HexToHash(rawTx))
	v, r, s := tx.RawSignatureValues()
	log.Println(common.BigToHash(v), common.BigToHash(r), common.BigToHash(s))
	log.Println(pending)
}
