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

func authenticateWithCookie(conf string, j *job) error {
	match := confRegExp.FindStringSubmatch(conf)
	if match == nil {
		return fmt.Errorf("conf format refused: %s", conf)
	}
	l := login{match[1], match[2]}
	b, err := json.Marshal(l)
	if err != nil {
		return err
	}
	resp, err := http.Post(match[3], "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusPartialContent {
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return fmt.Errorf("auth error: %s\n%v", resp.Status, err)
		}
		return fmt.Errorf("auth error: %s\n%s", resp.Status, string(b))
	}
	j.cookies = resp.Cookies()
	return nil
}

func authenticateWithToken(conf string, j *job) error {
	token := strings.Trim(conf, " ")
	j.headers = append(j.headers, [2]string{"Authorization", "Bearer " + token})
	return nil
}

func authenticate(conf string, j *job) error {
	if strings.HasPrefix(conf, "cookie ") {
		return authenticateWithCookie(conf[len("cookie "):], j)
	}
	if strings.HasPrefix(conf, "token ") {
		return authenticateWithToken(conf[len("token "):], j)
	}
	return fmt.Errorf("auth configuration not supported: %s", conf)
}
