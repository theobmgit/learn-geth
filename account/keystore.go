package main

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"log"
	"os"
)

func newKeystore(dir string) string {
	ks := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)
	pwd := "t4D0minic861"
	account, err := ks.NewAccount(pwd)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(account.URL)
	return account.Address.String()
}

func importKeystore(dir string, file string) {
	ks := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	pwd := "t4D0minic861"
	account, err2 := ks.Import(jsonBytes, pwd, pwd)
	if err2 != nil {
		log.Fatal(err2)
	}

	log.Println(account.URL.Path)

	if err = os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

func main() {
	newKeystore("./wallet")
}
