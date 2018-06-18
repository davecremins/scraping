package main

import (
	"sync"
)

type SiteData struct {
	url, jobsSelector, jobName string
}

func main() {
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
		go scrape(&wg, site.url, site.jobsSelector, site.jobName)
	}
	
	wg.Wait()
}