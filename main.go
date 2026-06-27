package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		err := hitURL(url)
		if err != nil {
			fmt.Println("error")
		}
	}
}

func hitURL(url string) error {
	fmt.Println("Checking ", url)
	resp, err := http.Get(url)
	fmt.Println(err)
	if resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
