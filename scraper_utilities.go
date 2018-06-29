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

func findName(i int, s *goquery.Selection, nameSelector string){
    name := s.Find(nameSelector).First().Text()
    color.Set(color.FgCyan, color.Bold)
    fmt.Printf("\t%d): %s \n", i+1, name)
    color.Unset()
}

func scrape(wg *sync.WaitGroup, url, container, itemName string){
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

    items := doc.Find(container)
    // Use handy standard colors
    color.Set(color.FgGreen, color.Bold)
    fmt.Printf("%d items found from %s \n", items.Length(), url)
    color.Unset()

    // Find specific info
    items.Each(func(i int, s *goquery.Selection) {
        findName(i, s, itemName)
    })

    fmt.Println()
}
