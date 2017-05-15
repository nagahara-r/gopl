// Copyright © 2017 Yuki Nagahara

package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

// Issue はGithub Issueの要素を示します。
type Issue struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	State string `json:"state"`
}

var (
	repository = "golang-testrepository"
	githubuser = "naga718"
)

// CreateIssue はIssueを作成します。
// ユーザ認証のため、パスワードを要求します。
func CreateIssue(item Issue) (err error) {
	url := "https://api.github.com/repos/" + githubuser + "/" + repository + "/issues"
	jsonByte, err := json.Marshal(item)
	if err != nil {
		return err
	}

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}

	req, err := GetGithubJSONRequest("POST", url, jsonByte, true)
	if err != nil {
		return err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	log.Println(resp.Status)

	return nil
}

// ReadIssue はIDに対してIssueを取得します。
func ReadIssue(id int) (issue Issue, err error) {
	no := strconv.Itoa(id)
	url := "https://api.github.com/repos/" + githubuser + "/" + repository + "/issues/" + no

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}
	req, err := GetGithubJSONRequest("GET", url, nil, false)
	if err != nil {
		return issue, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return issue, err
	}

	err = json.NewDecoder(resp.Body).Decode(&issue)
	return issue, err
}

// EditIssue はIssueを更新します。
func EditIssue(id int, issue Issue) (err error) {
	no := strconv.Itoa(id)
	url := "https://api.github.com/repos/" + githubuser + "/" + repository + "/issues/" + no

	jsonByte, err := json.Marshal(issue)
	if err != nil {
		return err
	}

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}
	req, err := GetGithubJSONRequest("PATCH", url, jsonByte, true)
	if err != nil {
		return err
	}

	resp, err := httpClient.Do(req)
	log.Println(resp.Status)

	return err
}

// CloseIssue はIssueをクローズします。
func CloseIssue(id int) (err error) {
	no := strconv.Itoa(id)
	url := "https://api.github.com/repos/" + githubuser + "/" + repository + "/issues/" + no

	json := "{\"state\":\"close\"}"

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}
	req, err := GetGithubJSONRequest("PATCH", url, []byte(json), true)
	if err != nil {
		return err
	}

	resp, err := httpClient.Do(req)
	log.Println(resp.Status)

	return err
}

// GetGithubJSONRequest はGithubに対するリクエストを生成します。
func GetGithubJSONRequest(method string, url string, body []byte, basicAuth bool) (req *http.Request, err error) {
	req, err = http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	if basicAuth {
		fmt.Printf("Password: ")
		var p string
		p, err = getPassword()

		req.SetBasicAuth("naga718", p)
	}

	return req, err
}

// BasicAuth用のパスワードを入力から取得します。
func getPassword() (password string, err error) {
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println("")
	if err != nil {
		return "", err
	}
	return string(bytePassword), nil
}

// SetGithubUser はGithubUserをセットします。
func SetGithubUser(user string) {
	githubuser = user
}

// SetRepository はリポジトリをセットします。
func SetRepository(repo string) {
	repository = repo
}
