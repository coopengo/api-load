package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

func do(j job, out chan<- job) {
	j.t = time.Now()
	defer func() {
		j.d = time.Since(j.t)
		out <- j
	}()
	req, err := http.NewRequest(j.method, j.url, bytes.NewReader(j.in))
	req.Header.Add("Content-Type", "application/json")
	for _, cookie := range j.cookies {
		req.AddCookie(cookie)
	}
	if err != nil {
		j.err = err.Error()
		return
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		j.err = err.Error()
		return
	}
	j.status = resp.StatusCode
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		j.err = err.Error()
		return
	}
	j.out = b
}

func worker(in <-chan job, out chan<- job) {
	for j := range in {
		do(j, out)
	}
	out <- job{}
}
