// Copyright © 2017 Yuki Nagahara

// 課題4-11：Github Issue を作成、読み出し、アップデート、クローズするコマンドラインツールです。
// 必要な場合はVimを開きます。
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/naga718/golang-practice/ch04/ex11/github"
)

func main() {
	u := flag.String("u", "", "[Required] Github User ID")
	r := flag.String("r", "", "[Required] Github Repository")

	c := flag.Bool("c", false, "Create Mode")
	re := flag.Bool("re", false, "Read Mode")
	e := flag.Bool("e", false, "Edit Mode")
	cl := flag.Bool("cl", false, "Close Mode")
	n := flag.Int("n", -1, "Issue Number (Ignore in Create Mode)")
	flag.Parse()

	if *u == "" || *r == "" {
		flag.Usage()
		return
	}

	github.SetGithubUser(*u)
	github.SetRepository(*r)

	if !*c && !*re && !*e && !*cl {
		flag.Usage()
		return
	}

	item := github.Issue{}

	if *c {
		fmt.Println("[Issue Create Mode]")
		fmt.Printf("Title: ")
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			log.Fatalf("Title Error")
		}

		item.Title = scanner.Text()

		fmt.Printf("Please Input Issue Description, Press Enter to Open Editor.")
		bufio.NewScanner(os.Stdin).Scan()

		bodyText, err := editBody("-Please Write your Issue-")
		if err != nil {
			log.Fatalf("%v", err.Error())
		}
		item.Body = bodyText

		err = github.CreateIssue(item)
		if err != nil {
			log.Fatalf("%v", err.Error())
		}
	} else if *re {
		fmt.Println("[Issue Read Mode]")
		if *n <= 0 {
			fmt.Println("Need: -n [IssueNumber]")
			return
		}

		issue, err := github.ReadIssue(*n)
		if err != nil {
			log.Fatalf("%v", err.Error())
		}

		fmt.Printf("No. %v\n", *n)
		fmt.Printf("Title: %v\n", issue.Title)
		fmt.Printf("Body: \n%v\n", issue.Body)
	} else if *e {
		fmt.Println("[Issue Update Mode]")
		if *n <= 0 {
			fmt.Println("Usage: $[program] -u -n [IssueNumber]")
			return
		}

		issue, err := github.ReadIssue(*n)
		if err != nil {
			log.Fatalf("%v", err.Error())
		}

		fmt.Printf("Please Edit Issue Title, Press Enter to Open Editor.")
		bufio.NewScanner(os.Stdin).Scan()
		outputText, err := editBody(issue.Title)
		item.Title = outputText

		fmt.Printf("Please Edit Issue Description, Press Enter to Open Editor.")
		bufio.NewScanner(os.Stdin).Scan()
		outputText, err = editBody(issue.Body)
		item.Body = outputText

		err = github.EditIssue(*n, item)
		if err != nil {
			log.Fatalf("%v", err.Error())
		}
	} else if *cl {
		fmt.Println("[Issue Close Mode]")
		if *n <= 0 {
			fmt.Println("Usage: $[program] -cl -n [IssueNumber]")
			return
		}

		github.CloseIssue(*n)
	}
}

func editBody(inputText string) (outputText string, err error) {
	file, err := os.Create("description.txt")
	if err != nil {
		log.Fatalf("%v", err.Error())
	}

	_, err = file.Write([]byte(inputText))
	if err != nil {
		return "", err
	}

	cmd := exec.Command("vim", "description.txt")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	defer exec.Command("rm", "description.txt").Run()
	if err != nil {
		return "", err
	}

	file, err = os.Open("description.txt")
	bodyByte, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(bodyByte), nil
}
