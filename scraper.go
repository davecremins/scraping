package main

import (
	"fmt"
  	"log"
  	"net/http"
	
	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Request the HTML page.
	res, err := http.Get("https://stackoverflow.com/jobs?sort=i&r=true&j=permanent")
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

	jobs := doc.Find("div[data-jobid]")

	found := jobs.Length()
	fmt.Printf("%d Jobs Found \n", found)

	// Find jobs
	jobs.Each(func(i int, s *goquery.Selection) {
		job := s.Find("h2 a").Text()
		fmt.Printf("Job %d: %s \n", i, job)
	})
}