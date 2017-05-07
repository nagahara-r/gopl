// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Issues prints a table of GitHub issues matching the search terms.

// Copyright © 2017 Yuki Nagahara
// 課題4-10： 一ヶ月未満、一年未満、一年以上の期間で分類し、結果を報告します。
package main

import (
	"fmt"
	"log"
	"os"

	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	monthly, yearly, moreThanYear := classifyIssues(result.Items)

	fmt.Println("--- Monthly Report Issues ---")
	printIssues(monthly)

	fmt.Println("--- Yearly Report Issues ---")
	printIssues(yearly)

	fmt.Println("--- More Than Year Report Issues ---")
	printIssues(moreThanYear)
}

func classifyIssues(items []*github.Issue) (monthly []*github.Issue, yearly []*github.Issue, moreThanYear []*github.Issue) {
	if items == nil {
		return
	}

	currentTime := time.Now()
	monthAgo := currentTime.AddDate(0, -1, 0)
	yearAgo := currentTime.AddDate(0, -12, 0)

	for _, item := range items {
		if monthAgo.Sub(item.CreatedAt) < 0 {
			monthly = append(monthly, item)
		} else if yearAgo.Sub(item.CreatedAt) < 0 {
			yearly = append(yearly, item)
		} else {
			moreThanYear = append(moreThanYear, item)
		}
	}

	return monthly, yearly, moreThanYear
}

func printIssues(items []*github.Issue) {
	for _, item := range items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

/*
//!+textoutput
19 issues:
--- Monthly Report Issues ---
#20206 markdryan encoding/base64: encoding is slow
--- Yearly Report Issues ---
#18227     LeGEC encoding/json: Decoder does not keep context for nested
#19526    Hunsin encoding/json: Wrong type in a field breaks decoding of
#16212 josharian encoding/json: do all reflect work before decoding
#17609 nathanjsw encoding/json: ambiguous fields are marshalled
#19469 chengzhic runtime: temporary object is not garbage collected
#15808 randall77 cmd/compile: loads/constants not lifted out of loop
#17244       adg proposal: decide policy for sub-repositories
#19029    cynecx runtime: fatal error: sweep increased allocation count
--- More Than Year Report Issues ---
#11046     kurin encoding/json: Decoder internally buffers full input
#15314    okdave proposal: some way to reject unknown fields in encoding
#12001 lukescott encoding/json: Marshaler/Unmarshaler not stream friendl
#5901        rsc encoding/json: allow override type marshaling
#8658  gopherbot encoding/json: use bufio
#14750 cyberphon encoding/json: parser ignores the case of member names
#7872  extempora encoding/json: Encoder internally buffers full output
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#7213  davechene cmd/compile: escape analysis oddity
#8717    dvyukov cmd/compile: random performance fluctuations after unre
//!-textoutput
*/
