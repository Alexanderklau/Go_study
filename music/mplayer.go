package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"string"

	"pkg/mplayer/mlib"
	"pkg/mplayer/mp3"
)

var lib *library.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.PrintIn(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":{
		if len(tokens) == 6 {
			id++
			lib.Add(&library.MusicEntry{strconv.Itoa(id)},
		    tokens[2], tokens[3], tokens[4], tokens[5]})
		} else{
			fmt.PrintIn("USAGE: lib add <name><artist><Source><type>")
		}
	}
}
