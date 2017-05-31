// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習7.9: printTracksをHTMLの表として曲を表示します。

// See page 187.

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

var trackList = template.Must(template.New("tracklist").Parse(`
<h1>{{len .}} Tracks</h1>
<table border=1>
<tr style='text-align: left'>
  <th><a href='./?s=title'>Title</a></th>
  <th><a href='./?s=artist'>Artist</a></th>
  <th><a href='./?s=album'>Album</a></th>
  <th><a href='./?s=year'>Year</a></th>
  <th><a href='./?s=length'>Length</a></th>
</tr>
{{range .}}
<tr valign="top">
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

func main() {
	listenAddress := "localhost:8000"

	handler := func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("s")

		result, err := printTracks(query)
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

func printTracks(query string) ([]byte, error) {
	sortTrack(query)

	var buffer bytes.Buffer
	if err := trackList.Execute(&buffer, tracks); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func sortTrack(query string) {
	if query == "title" {
		sort.Sort(byTitle(tracks))
	} else if query == "artist" {
		sort.Sort(byArtist(tracks))
	} else if query == "length" {
		sort.Sort(byLength(tracks))
	} else if query == "year" {
		sort.Sort(byYear(tracks))
	} else if query == "album" {
		sort.Sort(byAlbum(tracks))
	}
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//!+artistcode
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-artistcode

//!+yearcode
type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-yearcode

//!+titlecode
type byTitle []*Track

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-titlecode

//!+albumcode
type byAlbum []*Track

func (x byAlbum) Len() int           { return len(x) }
func (x byAlbum) Less(i, j int) bool { return x[i].Album < x[j].Album }
func (x byAlbum) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!+albumcode

//!+albumcode
type byLength []*Track

func (x byLength) Len() int           { return len(x) }
func (x byLength) Less(i, j int) bool { return x[i].Length < x[j].Length }
func (x byLength) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!+albumcode

func MultiTrackSort(first sort.Interface, second sort.Interface) {
	sort.Sort(second)
	sort.Sort(first)
}
