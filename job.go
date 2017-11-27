package main

import (
	"fmt"
	"time"
)

type job struct {
	id     int
	url    string
	method string
	req    []byte
	res    []byte
	status string
	err    string
	t      time.Time
	d      time.Duration
}

func (j *job) String() string {
	return fmt.Sprintf("%d %v %v", j.id, j.status, j.d)
}
