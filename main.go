package main

import (
	"fmt"
  	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

type SiteData struct {
	url, jobsSelector, jobName string
}

func main() {
	
	findJobName := func(i int, s *goquery.Selection, selector string){
		job := s.Find(selector).First().Text()
		color.Set(color.FgCyan, color.Bold)
		fmt.Printf("\tJob %d: %s \n", i+1, job)
		color.Unset()
	}
	
	
	sites := [] *SiteData{
		&SiteData{"https://weworkremotely.com/categories/remote-programming-jobs", "article ul li", "span.title"},
		&SiteData{"https://stackoverflow.com/jobs?r=true&j=permanent", "div[data-jobid]", "h2 a"},
		&SiteData{"https://remoteok.io/remote-dev-jobs", "tbody tr[id]", "a h2"},
		&SiteData{"https://www.workingnomads.co/jobs?category=development", "#jobs div.job", "a h2"},
		&SiteData{"https://jobspresso.co/remote-software-jobs/", "div.job_listings ul.job_listings li", "a"},
		
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