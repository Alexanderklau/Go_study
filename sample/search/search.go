package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run() {
	feeds, err := RetieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	results := make(chan *Result)
	var watiGroup sync.WaitGroup

	watiGroup.Add(len(feeds))

	for _, feed := range feeds {
		matchers, exists := matchers[feeds.Type]
		if !exists {
			matcher = matchers["deafault"]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			watiGroup.Done()
		}(matcher, feed)
	}

	go func() {
		watiGroup.Wait()
		close(results)
	}()

	Display(results)
}
