package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

const (
	baseURL = "https://www.sports-reference.com/cfb/players/%v-index.html"
)

func main() {
	playerCollector := colly.NewCollector()

	// Visits each player's individual page
	playerCollector.OnHTML(".section_content > p", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		e.Request.Visit(link)
	})

	// Visits the next page for the current letter in the index
	playerCollector.OnHTML(".next", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		e.Request.Visit(link)
	})

	playerCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// Iterates through each letter of the player index
	for i := 1; i < 27; i++ {
		link := fmt.Sprintf(baseURL, string('a'-1+i))
		playerCollector.Visit(link)
	}
}
