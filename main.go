package main

import (
	"flag"
	"fmt"
	"sync"
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
		url:     *url,
		method:  *method,
		headers: make([][2]string, 0, 1),
	}
	if err := authenticate(*auth, tpl); err != nil {
		panic(err)
	}

	errC := make(chan error, 1)
	todo := make(chan job, *buffer**concurrency)
	done := make(chan job, *buffer**concurrency)
	var workersWG sync.WaitGroup

	for i := uint(0); i < *concurrency; i++ {
		workersWG.Add(1)
		go worker(todo, done, &workersWG)
	}

	go prepare(todo, errC, tpl, path)

	go func() {
		err := <-errC
		panic(err)
	}()

	var loggerWG sync.WaitGroup
	loggerWG.Add(1)
	go log(done, &loggerWG)

	workersWG.Wait()
	close(done)
	loggerWG.Wait()
}
