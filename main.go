//go:generate abigen --sol Spawn.sol --pkg main --out Spawn.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

type issue struct {
	id   int
	desc string
}

var issues []issue

func main() {
	issues = append(issues, issue{5, "hello"})
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Issues ", issues)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
