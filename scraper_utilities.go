package main

import (
	"fmt"
  "log"
  "net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

type action func(i int, item *goquery.Selection)

func scrape(wg *sync.WaitGroup, url, container string, process action){
	defer wg.Done()
	
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
	  log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
	  log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
  
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
	  log.Fatal(err)
	}

	jobs := doc.Find(container)
	// Use handy standard colors
	color.Set(color.FgGreen, color.Bold)
	fmt.Printf("%d Jobs found from %s \n", jobs.Length(), url)
	color.Unset()

	// Find specific info
	jobs.Each(func(i int, s *goquery.Selection) {
		process(i, s)
	})
}