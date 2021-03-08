package main

import (
	"fmt"

	colly "github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector(
		// Visit only domains buff.163.com, steamcommunity.com
		colly.AllowedDomains("buff.163.com"),
	)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://buff.163.com/market/?game=csgo#tab=selling&page_num=1")
}
