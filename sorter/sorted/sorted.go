package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var infine *string = flag.String("i", "infine", "File contains values for sorting")
var outfile *string = flag.String("o", "outfine", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValunes(infine string) (values []int, err error) {
	file, err := os.Open(infine)
	if err != nil {
		fmt.Println("Failed to open the input file ", infine)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}

		str := string(line)

		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

func main() {
	flag.Parse()

	if infine != nil {
		fmt.Println("infine =", *infine, "outfine =", *outfile, "algorithm =",
			*algorithm)
	}
	values, err := readValunes(*infine)
	if err == nil {
		fmt.Println("Read values:", values)
	} else {
		fmt.Println(err)
	}
}
