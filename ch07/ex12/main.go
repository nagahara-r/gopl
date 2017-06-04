// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習7-12：データベースの list をHTMLの表として出力します。

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var itemList = template.Must(template.New("itemlist").Parse(`
<h1>{{len .}} Items</h1>
<table border=1>
<tr style='text-align: left'>
  <th>Item</th>
  <th>Price</th>
</tr>
{{range $key, $value := .}}
<tr valign="top">
  <td>{{$key}}</td>
  <td>{{$value}}</td>
</tr>
{{end}}
</table>
`))

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/show", db.show)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

var mut = new(sync.Mutex)

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := itemList.Execute(w, db); err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "db error = %v\n", err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

// create はdbに新たなアイテムを追加します。
// すでに存在するアイテムの場合は BadRequest を返します。
func (db database) create(w http.ResponseWriter, req *http.Request) {
	mut.Lock()
	defer mut.Unlock()
	item := req.URL.Query().Get("item")
	sprice := req.URL.Query().Get("price")

	uprice, err := strconv.ParseFloat(sprice, 32)
	if err != nil || uprice < 0 {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %q\n", sprice)
		return
	}

	if _, ok := db[item]; !ok {
		db[item] = dollars(uprice)
		fmt.Fprintf(w, "created: item = %s, price %s\n", item, db[item])
	} else {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item already exists: %q\n", item)
	}
}

// show は指定のアイテムを読み出します。
func (db database) show(w http.ResponseWriter, req *http.Request) {
	mut.Lock()
	defer mut.Unlock()
	item := req.URL.Query().Get("item")

	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

// update はアイテムを更新します。
// アイテムが存在しない場合はエラーを表示します。
func (db database) update(w http.ResponseWriter, req *http.Request) {
	mut.Lock()
	defer mut.Unlock()
	item := req.URL.Query().Get("item")
	sprice := req.URL.Query().Get("price")

	uprice, err := strconv.ParseFloat(sprice, 32)
	if err != nil || uprice < 0 {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %q\n", sprice)
		return
	}

	if price, ok := db[item]; ok {
		db[item] = dollars(uprice)
		fmt.Fprintf(w, "updated: item = %s, price %s -> %s\n", item, price, db[item])
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

// delete はアイテムを更新します。
// アイテムが存在しない場合はエラーを表示します。
func (db database) delete(w http.ResponseWriter, req *http.Request) {
	mut.Lock()
	defer mut.Unlock()
	item := req.URL.Query().Get("item")

	if price, ok := db[item]; ok {
		delete(db, item)
		fmt.Fprintf(w, "deleted: item = %s, price %s\n", item, price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
