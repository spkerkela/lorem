package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var wordCount int
	if len(os.Args) > 1 {
		wordCount, _ = strconv.Atoi(os.Args[1])
	}

	loremWords := strings.Fields("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	loremWordsCount := len(loremWords)
	if wordCount <= 0 {
		wordCount = loremWordsCount
	}
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
