package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
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
			match = strings.ReplaceAll(match, ".", "#")
			match = strings.ReplaceAll(match, "/", "#")
			splitBy(match, "#", list)
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

func printIfUnique(text string, list map[string]struct{}) {
	if _, ok := list[text]; !ok {
		fmt.Println(text)
		list[text] = struct{}{}
	}
}
