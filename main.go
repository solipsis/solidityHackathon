package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

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

func root(w http.ResponseWriter, r *http.Request) {
	issue := a{
		Issue{"hello", "there", "plz help", "https://why.lord.org"},
		64,
		true,
	}

	num := 20
	as := make([]a, 20)

	for i := 0; i < num; i++ {
		as[i] = issue
	}

	t, _ := template.ParseFiles("style/template/issue.html")
	t.Execute(w, as)
}

func main() {
	go updateIssues()

	fs := http.FileServer(http.Dir("style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))

	http.HandleFunc("/", root)
	http.HandleFunc("/add_issue", addIssue)
	http.HandleFunc("/monies", monies)
	http.ListenAndServe(":6060", nil)
}

func addIssue(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("style/template/form.html")
	t.Execute(w, nil)
}

func monies(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)

	//TODO(gdavis) We probably need to use the form data...

	address := "asdfasdf" //TODO(gdavis) Hook this up with something else.

	t, _ := template.ParseFiles("style/template/monies.html")

	t.Execute(w, address)
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

var issues []a

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
}
