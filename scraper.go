package main

import (
	"fmt"
	"log"

	// Colly is an open source WEB scrapping library
	"github.com/gocolly/colly"
)

func main() {
	// scrapping logic
	c := colly.NewCollector()
	c.Visit("https://scrapeme.live/shop/")

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		// printing all URLs associated with the a links in the page
		fmt.Println("%v", e.Attr("href"))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	// downloading the target HTML page
	c.Visit("https://scrapeme.live/shop/")

	fmt.Println("Hello World!")
}
