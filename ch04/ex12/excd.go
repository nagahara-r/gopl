// Excd はウェブコミックExcdを検索します。
// Copyright © 2017 Yuki Nagahara

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Excd は人気漫画 ExcdのJSONデータです
type Excd struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

// ParseAll は全ファイルをパースします。
func ParseAll() (excds []Excd) {
	for i := 0; i < 1000; i++ {
		filepath := "json/" + strconv.Itoa(i+1) + ".json"
		file, err := os.Open(filepath)
		if err != nil {
			log.Println("Warning: " + filepath + " open failed")
			log.Println(err.Error())
			continue
		}

		byteJSON, err := ioutil.ReadAll(file)
		file.Close()

		if err != nil {
			log.Println("Warning: " + filepath + " read failed")
			log.Println(err.Error())
			continue
		}

		excd, err := Parse(byteJSON)
		if err != nil {
			log.Println("Warning: " + filepath + " JSON parse failed")
			log.Println(err.Error())
			continue
		}
		excds = append(excds, excd)
	}

	return excds
}

// Parse はJSON Byteスライスをパースします。
func Parse(jsonByte []byte) (excd Excd, err error) {
	err = json.Unmarshal(jsonByte, &excd)
	return excd, err
}

// Search はExcd のタイトル、本文からクエリに一致するものを検索します。
func Search(excds []Excd, query string) (result []Excd) {
	for _, excd := range excds {
		if strings.Contains(excd.Title, query) || strings.Contains(excd.Transcript, query) {
			result = append(result, excd)
		}
	}

	return result
}

// Print はスライスからURLと内容を表示します。
func Print(excds []Excd) {
	fmt.Printf("%v storys\n\n", len(excds))

	for _, excd := range excds {
		fmt.Printf("No. %v\n", excd.Num)
		fmt.Printf("%v\n", excd.Img)
		fmt.Printf("%v\n\n\n", excd.Transcript)
	}
}

func main() {
	q := flag.String("q", "", "Search Query")
	flag.Parse()
	if *q == "" {
		fmt.Println("Usage: excd -q [Search Query]")
		return
	}

	excds := ParseAll()
	log.Println("Ready.")

	result := Search(excds, "Apple")
	Print(result)
}
