package main

import (
	"log"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type links struct {
	visited    []string
	notVisited []string
	mux        sync.Mutex
}

func exist(item string, links *links) (exists bool) {
	links.mux.Lock()
	defer links.mux.Unlock()

	for i := 0; i < len(links.visited); i++ {
		if links.visited[i] == item {
			return true
		}
	}

	for i := 0; i < len(links.notVisited); i++ {
		if links.notVisited[i] == item {
			return true
		}
	}

	return false
}

func spider(url, links *links) {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		linkTag := item
		link, _ := linkTag.Attr("href")

		if !strings.HasPrefix(link, url) {
			return
		}

		if exist(link, links) {
			return
		}

		if exist(link, links) {
			return
		}
		links.mux.Lock()
		links.notVisited = append(links.notVisited, link)
		links.mux.Unlock()
	})
}

func main() {

	url := "http://www.illesbalearsfilm.com"

	var links links
	links.notVisited = append(links.notVisited, url)
	spider(url, &links)
}
