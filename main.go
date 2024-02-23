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

	// Set User-Agent
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"

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
