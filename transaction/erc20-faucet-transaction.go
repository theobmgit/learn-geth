package main

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/ee09af9e74a243c4a7ee638fcf0a9d21")
	if err != nil {
		log.Fatalln(err)
	}

	//privateKey, _ := crypto.HexToECDSA("fd95304a7c242e89305064a3cdee4fdcd91762d10f54d4f0cd38a9d0b6b1e0bc")
	fromAddress := common.HexToAddress("0x48163ddc4d8149Dd544961b76B34b7987FAD3Ff5")
	toAddress := common.HexToAddress("0x101848D5C5bBca18E6b4431eEdF6B95E9ADF82FA")

	value := big.NewInt(0)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	gasTip, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	gasLimit, _ := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:    &toAddress,
		Value: value,
	})

	log.Println(nonce, gasPrice, gasTip, gasLimit)

	//chainId, _ := client.NetworkID(context.Background())
	//
	//tx := types.NewTx(&types.DynamicFeeTx{
	//	ChainID:   chainId,
	//	Nonce:     nonce,
	//	GasFeeCap: gasPrice,
	//	GasTipCap: gasTip,
	//	Gas:       gasLimit,
	//	To:        &toAddress,
	//	Value:     value,
	//})
	//
	//signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainId), privateKey)
	//if err != nil {
	//	log.Fatal("Sign tx", err)
	//}
	//
	//err = client.SendTransaction(context.Background(), signedTx)
	//if err != nil {
	//	log.Fatal("Send tx", err)
	//}
	//
	//log.Println(signedTx.Hash().Hex())
}
