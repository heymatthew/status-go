package main

import "testing"

func TestString(t *testing.T) {
	mockResponse := "200 AY-OK"
	mockGet := func() string {
		return mockResponse
	}

	testWebpage := webpage{url: "https://gitlab.com", getResponse: mockGet}

	if testWebpage.String() != "https://gitlab.com" {
		t.Errorf("String version of webpage doesn't translate")
	}
}
