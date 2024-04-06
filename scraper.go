package main

import (
	"fmt"
	"os"

	// Colly is an open source WEB scrapping library
	"github.com/gocolly/colly"
)

func main() {
	// scrapping logic
	args := os.Args
	url := args[1]
	c := colly.NewCollector()
	c.Visit("https://scrapeme.live/shop/")

	// whenever the collector is about to make a new request
	c.OnRequest(func(r *colly.Request) {
		// print the url of that request
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Blimey, an error occurred!:", e)
	})
	c.Visit(url)

	fmt.Println("Finish!")
}
