package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var confRegExp = regexp.MustCompile(`(\S+):(\S+)@(http\S+)`)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func authenticate(conf string, j *job) {
	if !strings.HasPrefix(conf, "cookie ") {
		panic(fmt.Errorf("auth configuration not supported: %s", conf))
	}
	conf = conf[len("cookie "):]
	match := confRegExp.FindStringSubmatch(conf)
	if match == nil {
		panic(fmt.Errorf("conf format refused: %s", conf))
	}
	l := login{match[1], match[2]}
	b, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(match[3], "application/json", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusPartialContent {
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			panic(fmt.Errorf("auth error: %s\n%v", resp.Status, err))
		} else {
			panic(fmt.Errorf("auth error: %s\n%s", resp.Status, string(b)))
		}
	}
	j.cookies = resp.Cookies()
}
