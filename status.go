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

// timedResponse returns the http status from url and ms passed since fetch
func (this webpage) timedResponse() string {
	start := time.Now()
	code, _ := this.responseCode()
	seconds := time.Since(start).Seconds()

	return fmt.Sprintf("%d %dms", code, int(seconds*1000))
}

// main starts the program, printing the http status and response time for 5 minutes
func main() {
	site := webpage{url: "https://gitlab.com", fetch: http.DefaultClient}
	_, _ = site.responseCode() // Hack, for some reason the first run is very slow

	start := time.Now()
	for time.Since(start).Minutes() < 5 {
		fmt.Println(site.timedResponse())
	}
}
