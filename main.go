package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	relDate, relVersion, relCommit string
)

func version() {
	fmt.Printf("Date    : %s\n", relDate)
	fmt.Printf("Version : %s\n", relVersion)
	fmt.Printf("Commit  : %s\n", relCommit)
}

func main() {
	// version
	v := flag.Bool("version", false, "print version informations")

	// func args
	auth := flag.String("auth", "cookie admin:admin@http://localhost:3000/auth/login", "api auth method")
	url := flag.String("url", "http://localhost:3000/contract", "api url")
	method := flag.String("method", "POST", "http method to process to api")

	// tech args
	concurrency := flag.Uint("c", 4, "number of concurrent requests")
	buffer := flag.Uint("b", 10, "buffer / channel size")

	flag.Parse()
	path := flag.Arg(0)

	if *v {
		version()
		return
	}

	if path == "" {
		flag.Usage()
		return
	}

	tpl := &job{
		url:    *url,
		method: *method,
	}
	authenticate(*auth, tpl)

	todo := make(chan job, *buffer**concurrency)
	done := make(chan job, *buffer**concurrency)

	for i := uint(0); i < *concurrency; i++ {
		go worker(todo, done)
	}

	go prepare(todo, tpl, path)

	wDone := uint(0)
	count, ok, ko := 0, 0, 0
	var dmin, dmax, sum time.Duration
	for {
		j := <-done
		if j.id == 0 {
			wDone++
			if wDone == *concurrency {
				close(done)
				break
			}
		} else {
			if j.Ok() {
				ok++
			} else {
				ko++
				fmt.Fprintln(os.Stderr, &j)
			}
			count++
			sum += j.d
			if j.d > dmax {
				dmax = j.d
			}
			if dmin == 0 || j.d < dmin {
				dmin = j.d
			}
		}
		fmt.Printf("\r%d Succeeded - %d Failed", ok, ko)
	}
	fmt.Printf("\nmin: %v - max: %s - avg: %v\n", dmin, dmax, sum/time.Duration(count))
}
