package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

func main() {
	list := make(map[string]struct{}) // store the output lines to check for dupes
	s := bufio.NewScanner(os.Stdin)
	r := regexp.MustCompile(`[a-zA-Z0-9\.\-\_\/]*`)
	for s.Scan() {
		process(s.Text(), list, r)
	}
	if s.Err() != nil {
		log.Printf("Error: %s\n", s.Err())
	}
}

func process(text string, list map[string]struct{}, r *regexp.Regexp) {
	for _, match := range r.FindAllString(text, -1) {
		if match != "" {
			if strings.Contains(match, ".") {
				splitBy(match, ".", list)
			}
			if strings.Contains(match, "/") {
				splitBy(match, "/", list)
			}
			printIfUnique(match, list)
		}
	}
}

func splitBy(text string, split string, list map[string]struct{}) {
	splitString := strings.Split(text, split)
	for _, match := range splitString {
		if match != "" {
			printIfUnique(match, list)
		}
	}
}

func removeSlashPrefix(text string) string {
	for {
		if strings.HasPrefix(text, "/") {
			_, i := utf8.DecodeRuneInString(text)
			text = text[i:]
			continue
		}
		break
	}
	return text
}

func printIfUnique(text string, list map[string]struct{}) {
	text = removeSlashPrefix(text)
	if _, ok := list[text]; !ok {
		fmt.Println(text)
		list[text] = struct{}{}
	}
}
