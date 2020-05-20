package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("please type url you wont to check")
		return
	}
	url := os.Args[1]
	fmt.Println(url)
	url = strings.Trim(url, "\"")
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("\n\n an error \n\n ", err)
		return
	}

	// if there is any re-direction happening behind the scene
	// the finalURL will be different
	// in this case, there will be a re-direction to https (SSL) version

	finalURL := resp.Request.URL.String()

	// Check if served with https
	if strings.HasPrefix(finalURL, "https") {
		fmt.Printf("%s is secured\n", finalURL)
		return
	}
	fmt.Printf("%s is not secured\n", finalURL)

}
