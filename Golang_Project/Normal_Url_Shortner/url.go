package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
)

var shortURL = make(map[string]string)

func main() {

	var input string
	var ch int
	var short_url string

	for {
		fmt.Println("\n----URL Shortener Program----")
		fmt.Print("1.Enter the URL\n2.Get the short URL\n3.Expand the URL\n4.Exit\n")
		fmt.Println("Enter your choice : ")
		fmt.Scan(&ch)

		switch ch {
		case 1:
			fmt.Println("Enter the URL : ")
			fmt.Scan(&input)
		case 2:
			short_url = shorten(input)
			fmt.Println("Shorten URL is : ", short_url)
			shortURL[short_url] = input
		case 3:

			original_url, found := expand(short_url)
			if found {
				fmt.Println("Expanded URL : ", original_url)
			} else {
				fmt.Println("Short URL not found.")
			}
		case 4:
			os.Exit(1)
		default:
			os.Exit(1)

		}

	}

}
// function to short the original url
func shorten(input string) string {
	hasher := sha1.New()
	hasher.Write([]byte(input))
	hasherValue := hasher.Sum(nil)
	//fmt.Println(hasherValue)
	hasherString := hex.EncodeToString(hasherValue)
	return hasherString[:12]
}
// function to expand the short url to original url
func expand(short_url string) (string, bool) {
	//original_url := hex.DecodeString(short_url)
	//fmt.Println("Expanded url is : ",original_url)

	longURL, found := shortURL[short_url]

	return longURL, found
}