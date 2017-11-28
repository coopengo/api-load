package main

import (
	"encoding/json"
	"io/ioutil"
)

func prepare(in chan job, temp *job, url string, method string, data string) {
	b, err := ioutil.ReadFile(data)
	if err != nil {
		panic(err)
	}
	objs := make([]map[string]interface{}, 0, 100)
	json.Unmarshal(b, &objs)
	for i, obj := range objs {
		j, err := json.Marshal(obj)
		if err != nil {
			panic(err)
		}
		in <- job{cookies: temp.cookies, id: i + 1, url: url, method: method, req: j}
	}
	close(in)
}
