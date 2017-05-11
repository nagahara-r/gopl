package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// OMDB の映画情報
type OMDB struct {
	Title  string
	Year   string
	Poster string
	Error  string
}

// GetOMDB はポスターのURLを取ってきます。
func GetOMDB(title string, year int) (omdb OMDB, err error) {
	t := "?t=" + url.QueryEscape(title)
	y := ""

	if year != 0 {
		y = "&y=" + strconv.Itoa(year)
	}
	omdbURL := "http://www.omdbapi.com/" + t + y

	json, err := Get(omdbURL)
	if err != nil {
		return omdb, err
	}

	omdb, err = Parse(json)
	if omdb.Error != "" {
		err = fmt.Errorf("No Movie Found")
	}

	log.Printf("Title: %v", omdb.Title)

	return omdb, err
}

// Parse はOMDBをパースします。
func Parse(data []byte) (omdb OMDB, err error) {
	err = json.Unmarshal(data, &omdb)
	return omdb, err
}

// Get はURLをByteスライスで入手します。
func Get(url string) (data []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	data, err = ioutil.ReadAll(resp.Body)
	return data, err
}

func main() {
	t := flag.String("t", "", "Movie Title")
	y := flag.Int("y", 0, "Movie Release Year")
	flag.Parse()

	if *t == "" {
		flag.Usage()
	}

	omdb, err := GetOMDB(*t, *y)
	if err != nil {
		log.Fatalf("%v", err.Error())
	}

	if omdb.Error != "" {
		log.Printf("%v", omdb.Error)
		return
	}
	posterData, err := Get(omdb.Poster)

	os.Stdout.Write(posterData)
}
