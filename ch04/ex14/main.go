// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習問題4-14: Githubへの一度の問い合わせでバグレポート、マイルストーン、ユーザ一覧を閲覧可能にします。

// See page 115.

// Issueshtml prints an HTML table of issues matching the search terms.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopl.io/ch4/github"

	"html/template"
)

//!+template

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table border=1>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>Milestone</th>
  <th>User</th>
  <th>Title / Description</th>
</tr>
{{range .Items}}
<tr valign="top">
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.Milestone.HTMLURL}}'>{{.Milestone.Title}}</a></td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><b><a href='{{.HTMLURL}}'>{{.Title}}</a></b></br>{{.Body}}</td>
</tr>
{{end}}
</table>
`))

func main() {
	listenAddress := "localhost:8000"

	handler := func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")

		result, err := webserverProc(query)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "400 bad request\n error = %v", err.Error())
		} else {
			w.Write(result)
		}
	}
	http.HandleFunc("/", handler)

	fmt.Println("Start Server: ", listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}

func webserverProc(query string) ([]byte, error) {
	queryArray := []string{query}
	result, err := SearchIssues(queryArray)
	if err != nil {
		return nil, err
	}

	for i := range result.Items {
		if result.Items[i].Milestone == nil {
			result.Items[i].Milestone = &Milestone{}
		}

		result.Items[i].Body = strings.Replace(result.Items[i].Body, "\n", "</br>", -1)
	}
	var buffer bytes.Buffer
	if err := issueList.Execute(&buffer, result); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(github.IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time     `json:"created_at"`
	Body      template.HTML // in Markdown format
	Milestone *Milestone
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Milestone struct {
	Title   string
	Number  int
	HTMLURL string `json:"html_url"`
}
