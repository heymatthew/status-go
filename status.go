// Package main contains methods to check the http status of a website.
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const DefaultUrl = "https://gitlab.com"

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

	defer resp.Body.Close()
	return resp.StatusCode, nil
}

// timedResponse returns the http status from url and ms passed since fetch
func (this webpage) timedResponse() string {
	start := time.Now()
	code, err := this.responseCode()
	milliseconds := int(time.Since(start).Seconds() * 1000)

	if err != nil {
		return fmt.Sprintf("%s %dms", err.Error(), milliseconds)
	}

	return fmt.Sprintf("%d %dms", code, milliseconds)
}

// getUrl returns the first runtime argument, or the default url
func getUrl() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return DefaultUrl
}

// main starts the program, printing the http status and response time for 5 minutes
func main() {
	site := webpage{url: getUrl(), fetch: http.DefaultClient}
	_, _ = site.responseCode() // Hack, for some reason the first run is very slow

	ticker := time.NewTicker(time.Second) // Flood protection
	go func() {
		for range ticker.C {
			fmt.Println(site.timedResponse())
		}
	}()

	time.Sleep(time.Minute * 5)
	ticker.Stop()
}
