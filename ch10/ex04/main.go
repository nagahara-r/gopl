// Copyright © 2017 Yuki Nagahara

// 練習10-4: 引数で指定したパッケージに推移的に依存している
// ワークスペース内のすべてのパッケージの集まりを報告するツールを作成します。

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Package struct {
	Imports []string
}

func main() {
	depends := make(map[string]bool)

	var p string
	flag.StringVar(&p, "p", "", "package name")
	flag.Parse()

	lsdata, err := execGoList(p)
	if err != nil {
		log.Fatalf("%v", err)
	}

	firstpack := new(Package)
	err = json.Unmarshal(lsdata, firstpack)
	if err != nil {
		log.Fatalf("json.Unmarshal(): %v", err)
	}

	depends = appendMap(depends, firstpack.Imports...)

	for _, p := range firstpack.Imports {
		lsdata, err = execGoList(p)
		if err != nil {
			log.Fatalf("%v", err)
		}

		pack := new(Package)
		err = json.Unmarshal(lsdata, pack)
		if err != nil {
			log.Fatalf("json.Unmarshal(): %v", err)
		}

		depends = appendMap(depends, pack.Imports...)
	}

	var imports []string
	for k := range depends {
		imports = append(imports, k)
	}

	fmt.Printf("package: %v\n", p)
	fmt.Printf("imports: \n%v\n", strings.Join(imports, ", "))
}

func appendMap(m map[string]bool, ss ...string) map[string]bool {
	for _, s := range ss {
		m[s] = true
	}

	return m
}

func execGoList(packagename string) (json []byte, err error) {
	if packagename == "" {
		return nil, fmt.Errorf("package can't be void")
	}

	lcmd := exec.Command("go", "list", "-json", packagename)
	w := new(bytes.Buffer)
	lcmd.Stdout = w

	err = lcmd.Start()
	if err != nil {
		return nil, fmt.Errorf("lcmd.Start(): %v", err)
	}
	err = lcmd.Wait()
	if err != nil {
		return nil, fmt.Errorf("lcmd.Wait(): %v", err)
	}

	return w.Bytes(), nil
}
