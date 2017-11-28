package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	res, err := http.Post(match[3], "application/json", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	j.cookies = res.Cookies()
}
