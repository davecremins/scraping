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
type SiteData struct {
	url, jobsSelector, jobName string
}

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
	fmt.Printf("%d Jobs Found \n", jobs.Length())
	color.Unset()

	// Find specific info
	jobs.Each(func(i int, s *goquery.Selection) {
		process(i, s)
	})
}

func main() {
	
	findJobName := func(i int, s *goquery.Selection, selector string){
		job := s.Find(selector).First().Text()
		color.Set(color.FgCyan, color.Bold)
		fmt.Printf("Job %d: %s \n", i+1, job)
		color.Unset()
	}
	
	
	sites := [] *SiteData{
		&SiteData{"https://weworkremotely.com/categories/remote-programming-jobs", "article ul li", "span.title"},
		&SiteData{"https://stackoverflow.com/jobs?r=true&j=permanent", "div[data-jobid]", "h2 a"},
		&SiteData{"https://remoteok.io/remote-dev-jobs", "tbody tr[id]", "a h2"},
	}
	
	var wg sync.WaitGroup
	wg.Add(len(sites))
	
	for _, site := range sites {
		go scrape(&wg, site.url, site.jobsSelector, func(s *SiteData) action {
			return func(i int, item *goquery.Selection) {
				findJobName(i, item, s.jobName)
			}
		}(site))
	}
	
	wg.Wait()
}