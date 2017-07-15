//go:generate abigen --sol Spawn.sol --pkg main --out Spawn.go
package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

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

	address, tx, spawner, err := DeploySpawn(auth, conn)
	if err != nil {
		log.Fatalf("Failed to deploy new contract %v", err)
	}
	fmt.Printf("Contract pending deploy 0x%x\n\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	fmt.Println("spawner: ", spawner)
	waitForBlock()

	session := &SpawnSession{
		Contract: spawner,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: big.NewInt(3141592),
		},
	}

	session.newIssue("CoolName", 5000)
	waitForBlock()
	session.newIssue("Issue222", 1984)
	waitForBlock()

	session.printIssues()

	issues = append(issues, issue{5, "hello"})
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Issues ", issues)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func waitForBlock() {
	time.Sleep(5000 * time.Millisecond)
}

func (session *SpawnSession) newIssue(name string, threshold int64) {
	issue, err := session.CreateIssue(name, big.NewInt(threshold))
	if err != nil {
		log.Fatalf("failed to make sub transfer :name:%V %v", name, err)
	}
	fmt.Println("Issue:", issue)
}

func (session *SpawnSession) fetchIssues() {

}

func (session *SpawnSession) printIssues() {
	fmt.Println("Printing issues")
	size, err := session.Size()
	if err != nil {
		log.Fatalf("Failed to retrieve index: %v", err)
	}
	fmt.Println("Size: ", size)
	for i := int64(0); i < size.Int64(); i++ {
		addr, err := session.AddressLUT(big.NewInt(i))
		if err != nil {
			log.Fatalf("Failed to retrieve address: %v", err)
		}
		issue, err := session.Issues(addr)
		if err != nil {
			log.Fatalf("Failed to retrieve issue from address: %v", err)
		}
		fmt.Printf("Name: %v Threshold: %v Index: %v", issue.Name, issue.Threshold, i)
	}
}
