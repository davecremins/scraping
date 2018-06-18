package main

import (
	"sync"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type SiteData struct {
	Url string `yaml:"site"`
	JobsSelector string `yaml:"job_listings"`
	JobName string `yaml:"job_title"`
}

type Config struct {
	Sites [] *SiteData `yaml:"site_info"`
 }

func main() {
	yamlFile, err := ioutil.ReadFile("config.yaml")

    if err != nil {
        panic(err)
	}
	
	var config Config

    err = yaml.Unmarshal(yamlFile, &config)
    if err != nil {
        panic(err)
    }

    var wg sync.WaitGroup
	wg.Add(len(config.Sites))
	
	for _, site := range config.Sites {
		go scrape(&wg, site.Url, site.JobsSelector, site.JobName)
	}
	
	wg.Wait()
}