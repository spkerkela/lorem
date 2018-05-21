package main

import (
	"bufio"
	"flag"
	"os"
	"strings"
)

func main() {
	loremWords := strings.Fields("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	loremWordsCount := len(loremWords)
	wordCountPtr := flag.Int("words", loremWordsCount, "number of words")
	flag.Parse()

	w := bufio.NewWriter(os.Stdout)
	for i := 0; i < *wordCountPtr; i++ {
		index := i % loremWordsCount
		w.WriteString(loremWords[index])
		if i != *wordCountPtr-1 {
			w.WriteString(" ")
		} else {
			w.WriteString("\n")
		}
	}

	w.Flush()

}
