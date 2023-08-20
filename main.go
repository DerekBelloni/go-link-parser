package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
	// Attribute string
}

func main() {
	fileName := "ex1.html"

	text, err := readHtmlFromFile(fileName)
	if err != nil {
		log.Fatal("Error reading html file: ", err)
	}

	links := parseTokens(text)

	fmt.Println("parsed links: ", links)
}

func parseTokens(text string) []Link {
	doc, err := html.Parse(strings.NewReader(text))
	if err !=  nil {
		log.Fatal("Error: ", err)
	}

	nodes := linkNodes(doc)
	var links []Link

	for _, n := range nodes{
		links = append(links, buildLink(n))
	}

	return links
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}

	return ret
}

func buildLink(n *html.Node) Link {
	var ret Link

	for _, a := range n.Attr {
		if a.Key == "href" {
			ret.Href = a.Val
		}
	}
	ret.Text = text(n)

	return ret
}

func text(n *html.Node) string {
	fmt.Println(n.Data)
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type == html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}

	return ret
}

func readHtmlFromFile(fileName string) (string, error) {
	htmlFile, err := os.ReadFile("ex2.html")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	return string(htmlFile), nil
}