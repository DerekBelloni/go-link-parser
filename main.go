package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/derekbelloni/go-link-parser/pkg/links"
)

func main() {
	fileName := "ex3.html"

	text, err := readHtmlFromFile(fileName)
	if err != nil {
		log.Fatal("Error reading html file: ", err)
	}

	r := strings.NewReader(text)

	links, err := links.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", links)
}

func readHtmlFromFile(fileName string) (string, error) {
	htmlFile, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	return string(htmlFile), nil
}

