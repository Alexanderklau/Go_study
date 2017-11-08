package search

import (
	"log"
)

type Result struct {
	Field   string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchTeam string) ([]*Result, error)
}


func Match(matcher Matcher, feed *Feed, searchTeam string, results chan<- *Result) {
    searchResults, err := matcher.Search(feed, searchTeam)
    if err != nil {
        log.Println("Errors print is:", err)
        return
    }

    for _, result := range searchResults {
        results <- result
    }

func Display(results chan *Result) {
    for result := range results {
        fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
    }
}

go func() {
    waitGroup.wait()
    close(result)
}()

