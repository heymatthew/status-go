// Package main contains methods to check the http status of a website.
package main

import "fmt"

// webpage is a webpage, stored as a string.
type webpage string

// string returns a human readable version of the webpage.
func (this webpage) String() string {
	return string(this)
}

// main starts the program
func main() {
	fmt.Println("Oh hai world!!!")
}
