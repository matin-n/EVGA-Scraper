package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"log"
)

func main() {

	link := "https://www.evga.com/products/ProductList.aspx?type=0&family=GeForce+30+Series+Family"

	c := colly.NewCollector(
		// http://go-colly.org/docs/introduction/configuration/
		// colly.AllowURLRevisit(),
		// colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36"),
	)

	extensions.RandomUserAgent(c)

	/*
		if p, err := proxy.RoundRobinProxySwitcher(
			"socks5://127.0.0.1:1337",
			"socks5://127.0.0.1:1338",
			"http://127.0.0.1:8080",
		); err == nil {
			c.SetProxyFunc(p)
		}

	*/

	// Find and scrape all
	c.OnHTML("div.pl-list-info", func(e *colly.HTMLElement) {

		fmt.Println(e.ChildTexts("div"))

		// TODO: this is no longer a reliable method to determine availability of product
		// Output if GPU found in stock
		//if e.ChildText("p.message.message-information") != "Out of Stock" {
		//	log.Println(
		//		e.ChildText("div.pl-list-pname"), // product name
		//		e.ChildText("p.message.message-information"), // out of stock or not
		//		e.Request.AbsoluteURL(e.ChildAttr("a", "href"))) // product URL
		//}
	})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL)
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Status Code:", r.StatusCode)
		log.Println("Request URL:", r.Request.URL, "failed with response:", string(r.Body), "\nError:", err)
	})

	c.Limit(&colly.LimitRule{
		// Parallelism: 10,
		// RandomDelay: 10 * time.Second,
	})

	c.Visit(link)

	// c.Wait()

}

