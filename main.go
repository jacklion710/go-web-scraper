package main

import (
    "fmt"
    "github.com/gocolly/colly/v2"
)

func main() {
    // Create a collector
    c := colly.NewCollector(
		colly.AllowedDomains("docs.cycling74.com"),
    )

    // On every a element which has href attribute call callback
    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        fmt.Printf("Link found: %q -> %s\n", e.Text, link)
    })

    // Start scraping
    fmt.Println("Starting to scrape...")
    c.Visit("https://docs.cycling74.com/max8/vignettes/rnbo_resources")

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
}
