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
	auth := flag.String("auth", "cookie admin:admin@http://localhost:3000/auth/login", "api auth method")
	url := flag.String("url", "http://localhost:3000/contract", "api url")
	method := flag.String("method", "POST", "http method to process to api")
	data := flag.String("data", "", "data file to process")

	flag.Parse()

	if *data == "" {
		flag.Usage()
		return
	}

	temp := &job{}
	authenticate(*auth, temp)

	in := make(chan job, *buffer**concurrency)
	out := make(chan job, *buffer**concurrency)

	for i := uint(0); i < *concurrency; i++ {
		go worker(in, out)
	}

	go prepare(in, temp, *url, *method, *data)

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
