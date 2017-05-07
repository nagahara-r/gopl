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

type Issue struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	State string `json:"state"`
}

const (
	githubuser = "naga718"
	repository = "golang-testrepository"
)

// Issueを作成します。
// ユーザ認証のため、パスワードを要求します。
func CreateIssue(item Issue) (err error) {
	url := "https://api.github.com/repos/" + githubuser + "/" + repository + "/issues"
	jsonByte, err := json.Marshal(item)
	if err != nil {
		return err
	}

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}

	req, err := GetGithubJsonRequest("POST", url, jsonByte, true)
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

// IDに対してIssueを取得します。
func ReadIssue(id int) (issue Issue, err error) {
	err = nil
	no := strconv.Itoa(id)
	url := "https://api.github.com/repos/" + githubuser + "/" + repository + "/issues/" + no

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}
	req, err := GetGithubJsonRequest("GET", url, nil, false)
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

// Issueを更新します。
func UpdateIssue(id int, issue Issue) (err error) {
	err = nil
	no := strconv.Itoa(id)
	url := "https://api.github.com/repos/" + githubuser + "/" + repository + "/issues/" + no

	jsonByte, err := json.Marshal(issue)

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}
	req, err := GetGithubJsonRequest("PATCH", url, jsonByte, true)
	if err != nil {
		return err
	}

	resp, err := httpClient.Do(req)
	log.Println(resp.Status)

	return err
}

// Issueをクローズします。
func CloseIssue(id int) (err error) {
	err = nil
	no := strconv.Itoa(id)
	url := "https://api.github.com/repos/" + githubuser + "/" + repository + "/issues/" + no

	json := "{\"state\":\"close\"}"

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}
	req, err := GetGithubJsonRequest("PATCH", url, []byte(json), true)
	if err != nil {
		return err
	}

	resp, err := httpClient.Do(req)
	log.Println(resp.Status)

	return err
}

// Githubに対するリクエストを生成します。
func GetGithubJsonRequest(method string, url string, body []byte, basicAuth bool) (req *http.Request, err error) {
	req, err = http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	if basicAuth {
		fmt.Printf("Password: ")
		password, err := password()
		if err != nil {
			return nil, err
		}

		req.SetBasicAuth("naga718", password)
	}

	return req, err
}

// BasicAuth用のパスワードを入力から取得します。
func password() (password string, err error) {
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println("")
	if err != nil {
		return "", err
	} else {
		return string(bytePassword), nil
	}
}
