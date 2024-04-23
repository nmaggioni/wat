package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func validateEntry(name []string, description []string, tags []string) bool {
	return (len(name) > 1 && len(name[1]) > 0) &&
		//(len(description) > 1 && len(description[1]) > 0) &&
		(len(tags) > 1 && len(tags[1]) > 0)
}

func main() {
	zshrc, err := os.Open(fmt.Sprintf("%s/.zshrc", os.Getenv("HOME")))
	if err != nil {
		panic(err)
	}
	defer zshrc.Close()

	nameRegex := regexp.MustCompile("^alias (.+?)=")                 // (?<=^alias ).+?(?==)
	descriptionRegex := regexp.MustCompile("#d:\\s?(.+?)(?: #t:|$)") // (?<= #d:).+?(?= #t:|$)
	tagsRegex := regexp.MustCompile("#t:\\s?(.+?)(?: #d:|$)")        // (?<= #t:).+?(?= #d:|$)

	scanner := bufio.NewScanner(zshrc)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "alias ") {
			nameMatches := nameRegex.FindStringSubmatch(line)
			descriptionMatches := descriptionRegex.FindStringSubmatch(line)
			tagsMatches := tagsRegex.FindStringSubmatch(line)

			if validateEntry(nameMatches, descriptionMatches, tagsMatches) {
				description := ""
				if len(descriptionMatches) > 1 {
					description = fmt.Sprintf("%s ", descriptionMatches[1])
				}
				tags := strings.Split(tagsMatches[1], ",")
				fmt.Printf(
					"%s: %s[%s]\n",
					strings.Trim(nameMatches[1], " "),
					strings.Trim(description, " "),
					strings.Join(tags, ","),
				)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
