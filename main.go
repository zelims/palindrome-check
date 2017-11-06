/**
 *
 * Palindrome Check
 * Author: Zach H.
 * Github: zelims
 * Date: 6 November, 2017
 *
 */

package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	word := flag.String("word", "", "Word to check for Palindrome")
	flag.Parse()
	if *word == "" {
		fmt.Printf("Usage: %s -word \"{WORD}\"\n", os.Args[0])
		return
	}
	fmt.Printf("Checking %s...%t\n", *word, checkPalindrome(*word))
}

func checkPalindrome(word string) bool {
	word = regexp.MustCompile("[^\\w]").ReplaceAllString(strings.ToLower(strings.Replace(word, " ", "", -1)), "")
	line := []rune(word)
	for i, str := range line {
		char := line[len(line)-(i+1):]
		if char[0] != str {
			return false
		}
	}
	return true
}