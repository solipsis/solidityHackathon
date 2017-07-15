//go:generate abigen --sol Spawn.sol --pkg main --out Spawn.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

type issue struct {
	id   int
	desc string
}

var issues []issue

const key = `{"address":"b5c423d9b40cb35ae2d70dc119a2db4d36a99ed7","crypto":{"cipher":"aes-128-ctr","ciphertext":"26876906681a87b4b90281d3ba85f6ed06cbbd037b9fdb7a95cafa2af95507d4","cipherparams":{"iv":"56ef1a75c04a9e7e1c54844ffe393bb3"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"d1c866674cff14a416acd2754aaea95188f0492b754e27daba1e3d0ba666c002"},"mac":"1d34f643b9a4a0b0ae44c4b5ac0f3d1152aa9d72f2c930ebe3fdb80df7b8af6a"},"id":"9b483bab-8006-4845-9248-48dbb734e068","version":3}`

func main() {

	conn, err := ethclient.Dial("/home/dave/testGeth/geth.ipc")
	if err != nil {
		log.Fatalf("Unable to connect to geth: %v\n", err)
	}
	fmt.Println("Connected successfully to geth:", conn)

	auth, err := bind.NewTransactor(strings.NewReader(key), "a")
	if err != nil {
		log.Fatalf("Unable to authenticate user: %v\n", err)
	}
	fmt.Println("Auth successfull", auth)

	fmt.Println("diff")
	address, tx, contract, err := DeploySpawn(auth)
	if err != nil {
		log.Fatalf("Unable to deploy contract: %v\n", err)
	}

	issues = append(issues, issue{5, "hello"})
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Issues ", issues)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
