// Package main contains methods to check the http status of a website.
package main

import "fmt"

// webpage is a webpage, stored as a string.
type webpage struct {
	url         string
	getResponse func() string
}

// string returns a human readable version of the webpage.
func (this webpage) String() string {
	return string(this.url)
}

// status returns the http response from a get request to this site
func (this webpage) status() string {
	return this.getResponse()
}

// main starts the program
func main() {
	fmt.Println("Oh hai world!!!")
}
