package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func prepareDir(todo chan<- job, tpl *job, path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for i, file := range files {
		b, err := ioutil.ReadFile(filepath.Join(path, file.Name()))
		if err != nil {
			panic(err)
		}
		j := *tpl
		j.id = i + 1
		j.in = b
		todo <- j
	}
}

func prepareFile(todo chan<- job, tpl *job, path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	_, err = dec.Token()
	if err != nil {
		panic(err)
	}
	i := 0
	for dec.More() {
		i++
		obj := make(map[string]interface{})
		err := dec.Decode(&obj)
		if err != nil {
			panic(err)
		}
		in, err := json.Marshal(obj)
		if err != nil {
			panic(err)
		}
		j := *tpl
		j.id = i
		j.in = in
		todo <- j
	}
}

func prepare(todo chan<- job, tpl *job, path string) {
	info, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	mode := info.Mode()

	if mode.IsDir() {
		prepareDir(todo, tpl, path)
	} else if mode.IsRegular() {
		prepareFile(todo, tpl, path)
	} else {
		panic(fmt.Errorf("can not manage path: %s", path))
	}
	close(todo)
}
