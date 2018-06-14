package main

import (
	"fmt"
  	"log"
  	"net/http"

	"github.com/PuerkitoBio/goquery"
)


func scrape(url, container string, process func(i int, item *goquery.Selection)){
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
	fmt.Printf("%d Jobs Found \n", jobs.Length())

	// Find specific info
	jobs.Each(func(i int, s *goquery.Selection) {
		process(i, s)
	})
}

func main() {
	scrape("https://stackoverflow.com/jobs?r=true&j=permanent", "div[data-jobid]", func(i int, s *goquery.Selection){
		job := s.Find("h2 a").Text()
		fmt.Printf("Job %d: %s \n", i, job)
	})
}