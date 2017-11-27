package main

import (
	"flag"
	"fmt"
)

func main() {
	// tech args
	concurrency := flag.Uint("c", 4, "number of concurrent requests")
	buffer := flag.Uint("b", 10, "buffer / channel size")
	// func args
	url := flag.String("url", "http://localhost", "api url to target")
	method := flag.String("method", "POST", "http method to process")
	data := flag.String("data", "", "data file to process")

	flag.Parse()

	if *data == "" {
		flag.Usage()
		return
	}

	in := make(chan job, *buffer**concurrency)
	out := make(chan job, *buffer**concurrency)

	for i := uint(0); i < *concurrency; i++ {
		go worker(in, out)
	}

	go prepare(in, *url, *method, *data)

	done := uint(0)
	for {
		r := <-out
		if r.id == 0 {
			done++
			if done == *concurrency {
				break
			}
		} else {
			fmt.Println(&r)
		}
	}
}
