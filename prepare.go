package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func prepare(in chan job, temp *job, url string, method string, data string) {
	info, err := os.Stat(data)
	if err != nil {
		panic(err)
	}
	mode := info.Mode()

	reqs := make([][]byte, 0, 100)
	if mode.IsDir() {
		files, err := ioutil.ReadDir(data)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			b, err := ioutil.ReadFile(filepath.Join(data, file.Name()))
			if err != nil {
				panic(err)
			}
			reqs = append(reqs, b)
		}
	} else if mode.IsRegular() {
		b, err := ioutil.ReadFile(data)
		if err != nil {
			panic(err)
		}
		objs := make([]map[string]interface{}, 0, 100)
		json.Unmarshal(b, &objs)
		for _, obj := range objs {
			req, err := json.Marshal(obj)
			if err != nil {
				panic(err)
			}
			reqs = append(reqs, req)
		}

	} else {
		panic(fmt.Errorf("can not manage path: %s", data))
	}
	for i, req := range reqs {
		in <- job{cookies: temp.cookies, id: i + 1, url: url, method: method, req: req}
	}
	close(in)
}
