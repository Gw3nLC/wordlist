package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// List of six letter words found by concatenating 2 words for the list
	var sixLetterWords []string

	path := os.Args[1]

	wordList, err := readLines(path)
	if err != nil {
		panic(err)
	}

	for i, word1 := range wordList {
		if len(string(word1)) >= 6 {
			// To long, ignored
		} else {
			for _, word2 := range wordList[i+1:] {
				if len(word1+word2) == 6 {
					// add the 6 letter word to the list
					sixLetterWords = append(sixLetterWords, word1+word2)
				}
			}
		}
	}

	// print the list of six letter words
	for _, sixLetterword := range sixLetterWords {
		fmt.Println(sixLetterword)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
