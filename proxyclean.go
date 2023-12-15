// proxyclean.go

package main

import (
    "math/rand"
    "flag"
    "fmt"
    "os"
    "bufio"
    "net/http"
    "sync"
    "time"
    //"math/rand"

    "golang.org/x/net/proxy"
)

// declare a shared vector of strings
var lines []string
var url = "http://www.google.com"
var wg sync.WaitGroup
var timeout_secs = 10 * time.Second
var verbose = false


func printtitle() {
    fmt.Println("\n",
"                                      __                \n",
"    ____  _________  _  ____  _______/ /__  ____ _____  \n",
"   / __ \\/ ___/ __ \\| |/_/ / / / ___/ / _ \\/ __ `/ __ \\ \n",
"  / /_/ / /  / /_/ />  </ /_/ / /__/ /  __/ /_/ / / / /\n",
" / .___/_/   \\____/_/|_|\\__, /\\___/_/\\___/\\__,_/_/ /_/ \n",
"/_/                    /____/                          \n",
"\n")

    fmt.Println("by darkmage")
    fmt.Println("https://www.github.com/mikedesu\n")
}

func usage() {
    fmt.Fprintln(os.Stderr, "Usage: proxyclean -f <filename> -t [threadcount]")
}

func readfile(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        usage()
        return
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        //fmt.Println(line)
        lines = append(lines, line)
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
}

func dorequest_socks5(proxya string) {
    // do request to url given
    dialer, err := proxy.SOCKS5("tcp", proxya, nil, proxy.Direct)
    if err != nil {
        wg.Done()
        return
    }
    // create client
    client := &http.Client{
        Timeout: timeout_secs,
        Transport: &http.Transport{
            Dial: dialer.Dial,
        },
    }
    req, err :=  client.Get(url)
    if err != nil {
        wg.Done()
        return
    }
    defer req.Body.Close()
    fmt.Println(proxya)
    //if verbose {
    //    fmt.Println(os.Stderr, "\033[33mSuccess\033[0m socks5://", proxya)
    //}
    wg.Done()
}

func main() {
    //printtitle()
    // pass file name as argument


    // get file name from command line
    filenamePtr := flag.String("f", "proxies.txt", "File name")
    threadcountPtr := flag.Int("t", 8, "Thread count")
    verbosePtr := flag.Bool("v", false, "Verbose output")
    flag.Parse()

    if *filenamePtr == "" {
        usage()
        // return an error to the OS
        os.Exit(1)
    }

    filename := *filenamePtr
    threadcount := *threadcountPtr
    verbose = *verbosePtr

    if threadcount <= 0 {
        fmt.Println(os.Stderr, "Thread count must be greater than 0")
        usage()
        os.Exit(1)
    }

    readfile(filename)
    // launch goroutine for proxies
    // we want to do these in batches of 20 or so
    // we also want to randomize the order
    // we can do this by pre-creating a list of indices
    // and then shuffling them

    indices := make([]int, len(lines))
    for i := range indices {
        indices[i] = i
    }
    // shuffle indices
    rand.Shuffle(len(indices), func(i, j int) { indices[i], indices[j] = indices[j], indices[i] })

    for j := 0; j+threadcount < len(lines); j+=threadcount {
        for i := j; i < j+threadcount; i++ {
            //var p = lines[i]
            var p = lines[indices[i]]
            //var pa = lines[rand.Intn(len(lines))]
            wg.Add(1)
            go dorequest_socks5(p)
        }
        // wait for goroutines to finish
        wg.Wait()
    }
}

