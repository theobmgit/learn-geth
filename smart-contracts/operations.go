package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	store "learn-geth/smart-contracts/contracts"
	"log"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/ee09af9e74a243c4a7ee638fcf0a9d21")
	if err != nil {
		log.Fatal(err)
	}
	//privateKey, err := crypto.HexToECDSA("fd95304a7c242e89305064a3cdee4fdcd91762d10f54d4f0cd38a9d0b6b1e0bc")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	//}
	//
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	address := common.HexToAddress("0xeBEC5C95f26Eb92d91e89D1A3B1083D00bc0D25f")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(instance.Version(nil))

	//key := [32]byte{}
	//value := [32]byte{}
	//copy(key[:], []byte("foo"))
	//copy(value[:], []byte("bar"))
	//
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//chainId, err := client.NetworkID(context.Background())
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.Value = big.NewInt(0)
	//auth.GasLimit = 300000
	//auth.GasPrice = gasPrice
	//
	//tx, err := instance.SetItem(auth, key, value)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("tx sent: %s", tx.Hash().Hex())

	var key [32]byte
	copy(key[:], "foo")
	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result[:]))
}
