package main

import (
	"fmt"
  	"log"
  	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
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
	fmt.Printf("%d Jobs Found \n", jobs.Length())

	// Find specific info
	jobs.Each(func(i int, s *goquery.Selection) {
		process(i, s)
	})
}

func main() {
	
	findJobName := func(i int, s *goquery.Selection, selector string){
		job := s.Find(selector).Text()
		fmt.Printf("Job %d: %s \n", i, job)
	}
	
	var wg sync.WaitGroup
	wg.Add(3)
        
	go scrape(&wg, "https://weworkremotely.com/categories/remote-programming-jobs", "article ul li", func(i int, s *goquery.Selection){
		findJobName(i, s, "span.title")
	})
	
	
	go scrape(&wg, "https://stackoverflow.com/jobs?r=true&j=permanent", "div[data-jobid]", func(i int, s *goquery.Selection){
		findJobName(i, s, "h2 a")
	})

	go scrape(&wg, "https://remoteok.io/remote-dev-jobs", "tbody tr[id]", func(i int, s *goquery.Selection){
		findJobName(i, s, "a h2")
	})
	
	wg.Wait()
}