package main

import "strings"

/*
The Challenge

Count the number of words provided. For this problem, a word is any sequence of
lowercase letters.

For example:
hello is a word, but so are non-English “words” like bbaabbb.

# Input

The input is one line of text, consisting of lowercase letters and spaces.
There is exactly one space between each pair of words, and there are no spaces
before the ﬁrst word or after the last word.
The maximum length of the line is 80 characters.

# Output

Output the number of words in the input line.
*/

// WordCounter counts the number of words
// A word is any sequence of lowercase letters
func WordCounter(input string) int {
	if input == "" {
		return 0
	}
	words := strings.Split(input, " ")
	return len(words)
}

// WordCounterNoSplit counts the number of words
// A word is any sequence of lowercase letters
// This function does not use the strings.Split function
func WordCounterNoSplit(input string) int {
	if input == "" {
		return 0
	}
	count := 1
	for _, char := range input {
		if char == ' ' {
			count++
		}
	}
	return count
}
