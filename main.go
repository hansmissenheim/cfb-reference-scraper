package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

const (
	baseURL = "https://www.sports-reference.com/cfb/players/%v-index.html"
)

type player struct {
	name string
}

func main() {
	playerCollector := colly.NewCollector()
	infoCollector := colly.NewCollector()

	// Visits each player's individual page
	playerCollector.OnHTML(".section_content > p", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		infoCollector.Visit(link)
	})

	// Visits the next page for the current letter in the index
	playerCollector.OnHTML(".next", func(e *colly.HTMLElement) {
		playerURL := e.Attr("href")
		e.Request.Visit(playerURL)
	})

	playerCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	infoCollector.OnHTML("div#info", func(e *colly.HTMLElement) {
		tempPlayer := player{}
		e.ForEach("h1 > span", func(_ int, el *colly.HTMLElement) {
			tempPlayer.name = el.Text
			fmt.Println(tempPlayer)
		})
	})

	infoCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Info visiting", r.URL)
	})

	// Iterates through each letter of the player index
	for i := 1; i < 27; i++ {
		link := fmt.Sprintf(baseURL, string('a'-1+i))
		playerCollector.Visit(link)
	}
	playerCollector.Visit("https://www.sports-reference.com/cfb/players/a-index-2.html")
}
