package main

import (
	"fmt"
	"net/http"
	"time"
)

type job struct {
	id      int
	url     string
	method  string
	err     string
	cookies []*http.Cookie
	in      []byte
	out     []byte
	status  int
	t       time.Time
	d       time.Duration
}

func (j *job) Ok() bool {
	return j.err == "" &&
		j.status >= http.StatusOK &&
		j.status <= http.StatusPartialContent
}

func (j *job) String() string {
	if j.Ok() {
		return fmt.Sprintf("%d %d %v", j.id, j.status, j.d)
	}
	err := j.err
	if err == "" {
		err = string(j.out)
	}
	return fmt.Sprintf("%d %d %v", j.id, j.status, err)
}
