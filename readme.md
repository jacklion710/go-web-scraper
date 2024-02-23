# Go Web Scraper

This project is a simple web scraper written in Go, using the Colly framework. It's designed to scrape links from web pages, showcasing the basic capability of Go in handling web scraping tasks.

## Features

- Scrapes all links from a given webpage.
- Filters links based on the domain to respect the website's boundaries.
- Prints the found links to the console, along with the text of the link.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.15 or later is recommended)
- Git

### Installing

First, clone the repository to your local machine:

```bash
git clone https://github.com/yourusername/go-web-scraper.git
```

Navigate into the project directory:

```bash
cd go-web-scraper
```

Install the project dependencies:

```bash
go mod tidy
```

## Running the Scraper

To run the scraper, simply execute:

```bash
go run main.go
```

## Built With

* [Go](https://go.dev) - The Go programming language
* [Colly](https://go-colly.org) - The scraping framework used