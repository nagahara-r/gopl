// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習8.9 rootディレクトリのそれぞれに対して個別の合計を計算し、定期的に表示するduの実装

// See page 250.

// The du3 command computes the disk usage of the files in a directory.
package main

// The du3 variant traverses all directories in parallel.
// It uses a concurrency-limiting counting semaphore
// to avoid opening too many files at once.

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//var vFlag = flag.Bool("v", false, "show verbose progress messages")
var tFlag = flag.Int64("t", 10, "update time (seconds)")

type dirInfo struct {
	index  int
	path   string
	nfiles int64
	nbytes int64
}

//!+
func main() {
	// ...determine roots...

	//!-
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	start(roots, *tFlag)
}

func obtainDirInfo(i int, root string, ch chan dirInfo) {
	//!+
	// Traverse each root of the file tree in parallel.
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	//for _, root := range roots {
	n.Add(1)
	go walkDir(root, &n, fileSizes)
	//}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	//!-

	// Print the results periodically.
	//var tick <-chan time.Time
	// if *vFlag {
	// 	tick = time.Tick(500 * time.Millisecond)
	// }
	var dir dirInfo
	dir.path = root
	dir.index = i
	//var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			dir.nfiles++
			dir.nbytes += size
			//case <-tick:
			//printDiskUsage(nfiles, nbytes)

		}
	}

	//printDiskUsage(nfiles, nbytes) // final totals
	ch <- dir

	//!+
	// ...select loop...
}

func start(roots []string, seconds int64) {
	// 周期的実行
	var tick <-chan time.Time
	tick = time.Tick(time.Second * time.Duration(seconds))
	routine(roots)

	for {
		select {
		case <-tick:
			routine(roots)
		}
	}
}

//!-

func routine(roots []string) {
	ch := make(chan dirInfo)
	dirs := make([]dirInfo, len(roots))
	for i, root := range roots {
		go obtainDirInfo(i, root, ch)
	}
	for i := range roots {
		dirs[i] = <-ch
	}

	for _, dir := range dirs {
		fmt.Printf("%v\n", dir.path)
		printDiskUsage(dir.nfiles, dir.nbytes)
		fmt.Print("\n")
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
//!+walkDir
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

//!-walkDir

//!+sema
// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...
	//!-sema

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
