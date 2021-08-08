package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

func main() {
	list := make(map[string]struct{}) // store the output lines to check for dupes
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		process(s.Text(), &list)
	}
}

func process(text string, list *map[string]struct{}) {
	r := regexp.MustCompile(`[a-zA-Z0-9\.\-\_\/]*`)
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

func splitBy(text string, split string, list *map[string]struct{}) {
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

func printIfUnique(text string, listPtr *map[string]struct{}) {
	text = removeSlashPrefix(text)
	list := *listPtr
	if _, ok := list[text]; !ok {
		fmt.Println(text)
		list[text] = struct{}{}
	}
}
