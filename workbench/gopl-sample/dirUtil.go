/************************************************
  > File Name: dirUtil.go
  > Author: Philo
  > Mail: 
  > Created Time: Tue 04 Feb 2020 08:53:12 AM CST
*************************************************/
package main

import (
    "fmt"
    "flag"
    "io/ioutil"
    "os"
    "path/filepath"
    "time"
    "sync"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
    flag.Parse()
    roots := flag.Args()
    if len(roots) == 0 {
        roots = []string{"."}
    }

    fileSizes := make(chan int64)

    /*go func() {
        for _, root := range roots {
            walkDir(root, fileSizes)
        }
        close(fileSizes)
    }()*/
    var n sync.WaitGroup
    for _, root := range roots {
        n.Add(1)
        go walkDir(root, &n, fileSizes)
    }
    go func() {
        n.Wait()
        close(fileSizes)
    }()

    /*var nfiles, nbytes int64
    for size := range fileSizes {
        nfiles++
        nbytes += size
    }*/

    var tick <-chan time.Time
    if *verbose {
        tick = time.Tick(500 * time.Millisecond)
    }
    var nfiles, nbytes int64
    loop:
    for {
        select {
        case size, ok := <-fileSizes:
            if ! ok {
                break loop
            }
            nfiles++
            nbytes += size
        case <-tick:
            printDiskUsage(nfiles, nbytes)
        }
    }

    printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
    fmt.Printf("%d files %.1f GB \n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
    defer n.Done()
    for _, entry := range dirents(dir) {
        //fmt.Printf("dir/file[%s] ...\n", entry.Name())
        if entry.IsDir() {
            n.Add(1)
            subdir := filepath.Join(dir, entry.Name())
            go walkDir(subdir, n, fileSizes)
        } else {
            fileSizes <- entry.Size()
        }
    }
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
    sema <- struct{}{}
    defer func() { <-sema }()
    entries, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Fprintf(os.Stderr, "du1: %v\n", err)
        return nil
    }
    return entries
}
