//go:generate abigen --sol Spawn.sol --pkg main --out Spawn.go
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type issue struct {
	id        string
	threshold int
}

var issues []issue

var conn *ethclient.Client
var auth *bind.TransactOpts
var session *SpawnSession

var tracked = make(map[string]common.Address)

const key = `{"address":"b5c423d9b40cb35ae2d70dc119a2db4d36a99ed7","crypto":{"cipher":"aes-128-ctr","ciphertext":"26876906681a87b4b90281d3ba85f6ed06cbbd037b9fdb7a95cafa2af95507d4","cipherparams":{"iv":"56ef1a75c04a9e7e1c54844ffe393bb3"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"d1c866674cff14a416acd2754aaea95188f0492b754e27daba1e3d0ba666c002"},"mac":"1d34f643b9a4a0b0ae44c4b5ac0f3d1152aa9d72f2c930ebe3fdb80df7b8af6a"},"id":"9b483bab-8006-4845-9248-48dbb734e068","version":3}`

func main() {

	conn, err := ethclient.Dial("/home/dave/testGeth/geth.ipc")
	if err != nil {
		log.Fatalf("Unable to connect to geth: %v\n", err)
	}
	fmt.Println("Connected successfully to geth:", conn)

	auth, err = bind.NewTransactor(strings.NewReader(key), "a")
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

	session = &SpawnSession{
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

	session.newIssue("https://github.com/ethereum/go-ethereum/issues/14803", 5000)
	waitForBlock()
	session.newIssue("https://github.com/solipsis/solidityHackathon/issues/7", 1984)
	waitForBlock()
	session.printIssues()
	session.fetchIssues(conn)

	issues = append(issues, issue{"www", 4})
	//	pollIssue("blah")

	//go updateIssues()

	fs := http.FileServer(http.Dir("style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))

	http.HandleFunc("/", root)
	http.HandleFunc("/add_issue", addIssue)
	http.HandleFunc("/monies", monies)
	http.ListenAndServe(":6060", nil)
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

func (session *SpawnSession) fetchIssues(conn *ethclient.Client) {
	size, err := session.Size()
	m := make(map[string]int)
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

		tracked[issue.Id] = addr

		transfer, err := NewTransfer(addr, conn)
		if err != nil {
			log.Fatalf("Failed to instantiate a Token contract: %v", err)
		}
		fmt.Println("ADDDRRESSS", addr)
		funded, err := transfer.Funded(&session.CallOpts)
		if err != nil {
			log.Fatalf("Failed to retrieve token name: %v", err)
		}

		if funded {
			m[issue.Id] += int(issue.Threshold.Int64())
		}

	}
	fmt.Println("MMMM: ", m)

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
		fmt.Printf("Name: %v Threshold: %v Index: %v", issue.Id, issue.Threshold, i)
	}

}

func pollIssue(url string) {
	_, err := http.Get("https://github.com/solipsis/solidityHackathon/issues/3")
	if err != nil {
		fmt.Printf("Error fetching issue %v\n", err)
	}
	//fmt.Println(resp.Body.)
}

type a struct {
	Issue  Issue
	Reward int
	Open   bool
}

type Issue struct {
	Name        string
	Project     string
	Description string
	URL         string
}

var as []a

func init() {

	issue := a{
		Issue{"hello", "there", "plz help", "https://why.lord.org"},
		64,
		true,
	}
	as = make([]a, 20)
	num := 20

	for i := 0; i < num; i++ {
		as[i] = issue
	}

}

func root(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("style/template/issue.html")
	t.Execute(w, as)
}

func addIssue(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("style/template/form.html")
	t.Execute(w, nil)
}

func monies(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)

	//session.newIssue(r.Form["issue"], r.Form["amount"])
	session.newIssue("https://github.com/ethereum/go-ethereum/issues/14803", 9999)
	address := common.Address{}
	//TODO(gdavis) We probably need to use the form data...
	for _, v := range tracked {
		address = v
	}
	//TODO(gdavis) Hook this up with something else.

	t, _ := template.ParseFiles("style/template/monies.html")

	as = append(as, a{Issue{"14803 https://github.com/ethereum/go-ethereum/issues/14803", "Add docs", "Description", "https://github.com/ethereum/go-ethereum/issues/14803"}, 9999, true})
	t.Execute(w, address.String())
}

func (a *a) isClosed() bool {

	if a.Issue.URL == "" {
		return false
	}

	ur := a.Issue.URL

	u, err := url.Parse(ur)
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Split(u.Path, "/")
	user := s[1]
	repo := s[2]
	number := s[4]

	resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s", user, repo, number))
	if err != nil {
		log.Fatal(err)
	}
	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	var x map[string]interface{}
	if err := json.Unmarshal(f, &x); err != nil {
		log.Fatal(err)
	}
	if _, ok := x["state"]; ok {
		if strings.ToLower(x["state"].(string)) == "closed" {
			a.Open = false
			return true
		}
	}

	a.Open = true
	return false
}

/*
var issues []

func updateIssues() {
	for {
		log.Println("waiting")
		time.Sleep(5 * time.Second)

		log.Println("reading issues")
		var wg sync.WaitGroup
		for _, issue := range issues {
			issue := issue
			go func() {
				wg.Add(1)
				if issue.isClosed() {
					// check things....
				}
			}()
		}
		wg.Wait()
	}
*/
