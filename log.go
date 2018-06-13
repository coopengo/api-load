package main

import (
	"fmt"
	"sync"
	"time"
)

func log(done <-chan job, wg *sync.WaitGroup) {
	count, ok, ko := 0, 0, 0
	fails := make([]job, 0, 0)
	var dmin, dmax, sum time.Duration
	for j := range done {
		if j.Ok() {
			ok++
		} else {
			ko++
			fails = append(fails, j)
		}
		count++
		sum += j.d
		if j.d > dmax {
			dmax = j.d
		}
		if dmin == 0 || j.d < dmin {
			dmin = j.d
		}
		fmt.Printf("\r%d Succeeded - %d Failed", ok, ko)
	}
	fmt.Printf("\nmin: %v - max: %s - avg: %v\n", dmin, dmax, sum/time.Duration(count))
	fmt.Println("\nFailed jobs")
	for _, j := range fails {
		fmt.Println(&j)
	}
	wg.Done()
}
