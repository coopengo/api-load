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
	cookies []*http.Cookie
	req     []byte
	res     []byte
	status  string
	err     string
	t       time.Time
	d       time.Duration
}

func (j *job) String() string {
	return fmt.Sprintf("%d %v %v", j.id, j.status, j.d)
}
