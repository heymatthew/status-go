// Package main contains methods to check the http status of a website.
package main

import (
	"fmt"
	"net/http"
	"time"
)

// webpage is a response time checker
type webpage struct {
	url   string
	fetch httpHead
}

// httpHead is an abstraction over the http.Head
type httpHead interface {
	Head(string) (resp *http.Response, err error)
}

// responseCode checks webpage.url and reutrns it's http resonse
func (this webpage) responseCode() (int, error) {
	resp, err := this.fetch.Head(this.url)
	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}

// main starts the program
func main() {
	site := webpage{url: "https://gitlab.com", fetch: http.DefaultClient}

	start := time.Now()
	code, _ := site.responseCode()
	seconds := time.Since(start).Seconds()

	timeMillis := fmt.Sprintf("%dms", int(seconds*1000))
	fmt.Println(code, timeMillis)
}
