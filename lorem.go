package main

import (
	"bufio"
	"flag"
	"os"
	"strings"
)

var wordCount int
var loremWords []string

func init() {
	loremWords = strings.Fields("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	loremWordsCount := len(loremWords)
	flag.IntVar(&wordCount, "words", loremWordsCount, "number of words")
}

func main() {
	flag.Parse()
	printLoremIpsum()
}

func printLoremIpsum() {
	loremWordsCount := len(loremWords)
	w := bufio.NewWriter(os.Stdout)
	for i := 0; i < wordCount; i++ {
		index := i % loremWordsCount
		w.WriteString(loremWords[index])
		if i != wordCount-1 {
			w.WriteString(" ")
		} else {
			w.WriteString("\n")
		}
	}

	w.Flush()

}
