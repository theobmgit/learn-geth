package main

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/ee09af9e74a243c4a7ee638fcf0a9d21")
	if err != nil {
		log.Fatalln(err)
	}

	privateKey, _ := crypto.HexToECDSA("fd95304a7c242e89305064a3cdee4fdcd91762d10f54d4f0cd38a9d0b6b1e0bc")
	fromAddress := common.HexToAddress("0x48163ddc4d8149Dd544961b76B34b7987FAD3Ff5")
	toAddress := common.HexToAddress("0x8c9c003fC635465b60bFd27f370E39CBC6AD770A")

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

	tokenAddress := common.HexToAddress("0x101848D5C5bBca18E6b4431eEdF6B95E9ADF82FA")
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodId := hash.Sum(nil)[:4]

	paddedToAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	amount := new(big.Int)
	amount.SetString("10000000000000000", 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodId...)
	data = append(data, paddedToAddress...)
	data = append(data, paddedAmount...)

	gasLimit, _ := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	})

	log.Println(nonce, gasPrice, gasTip, gasLimit)

	chainId, _ := client.NetworkID(context.Background())

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce,
		GasFeeCap: gasPrice,
		GasTipCap: gasTip,
		Gas:       gasLimit,
		To:        &toAddress,
		Value:     value,
		Data:      data,
	})

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainId), privateKey)
	if err != nil {
		log.Fatal("Sign tx", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("Send tx", err)
	}

	log.Println(signedTx.Hash().Hex())
}
