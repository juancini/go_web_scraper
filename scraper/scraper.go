package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// Colly is an open source WEB scrapping library
	"github.com/gocolly/colly"
	// SQLite drivers for database
	_ "github.com/mattn/go-sqlite3"
)

type PokemonProduct struct {
	url, image, name, price string
}

func main() {
	// input args
	args := os.Args
	url := args[1]

	// initializing sql database that will save the scraped data
	db, err := sql.Open("sqlite3", "scraping.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a table to store the scraped data
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS scraped_data (
            url TEXT,
			image TEXT,
			name TEXT,
			price TEXT
        )
    `)
	if err != nil {
		log.Fatal(err)
	}

	// initializing the slice of structs that will contain the scraped data
	var pokemonProducts []PokemonProduct

	// new colly collector
	c := colly.NewCollector()

	// scrapping logic
	// whenever the collector is about to make a new request
	c.OnRequest(func(r *colly.Request) {
		// print the url of that request
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		// initializing a new PokemonProduct instance
		pokemonProduct := PokemonProduct{}
		// scraping the data of interest
		pokemonProduct.url = e.ChildAttr("a", "href")
		pokemonProduct.image = e.ChildAttr("img", "src")
		pokemonProduct.name = e.ChildText("h2")
		pokemonProduct.price = e.ChildText(".price")

		// adding the product instance with scraped data to the list of products
		pokemonProducts = append(pokemonProducts, pokemonProduct)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Blimey, an error occurred!:", e)
	})

	// go to site of URL
	c.Visit(url)

	// writing each Pokemon to Database
	for _, pokemonProduct := range pokemonProducts {
		fmt.Println("Saving data...")
		// Insert the scraped data into the database
		_, err = db.Exec("INSERT INTO scraped_data (url, image, name, price) VALUES (?, ?, ?, ?)", pokemonProduct.url, pokemonProduct.image,
			pokemonProduct.name, pokemonProduct.price)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Finish!")
}
