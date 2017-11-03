package search

import (
	"log"
	"sync"
)

var matchers = make(string | Matcher)

func Run() {
	feeds, err := RetieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	results := make(chan *Result)
	var watiGroup sync.WaitGroup
}
