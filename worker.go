package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

func do(j job, out chan job) {
	j.t = time.Now()
	defer func() {
		j.d = time.Since(j.t)
		out <- j
	}()
	req, err := http.NewRequest(j.method, j.url, bytes.NewReader(j.req))
	req.Header["Content-Type"] = []string{"application/json"}
	if err != nil {
		j.err = err.Error()
		return
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		j.err = err.Error()
		return
	}
	j.status = res.Status
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		j.err = err.Error()
		return
	}
	j.res = b
}

func worker(in chan job, out chan job) {
	for j := range in {
		do(j, out)
	}
	out <- job{}
}
