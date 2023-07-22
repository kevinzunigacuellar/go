package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/net/html"
)

const rawHtml = `
<!DOCTYPE html>
<html>
	<body>
		<h1> one two three four five is an example</h1>
		<p>Some more words to count</p>
		<img src="https://golang.org/doc/gopher/frontpage.png" />
	</body>
</html>
`

func visitNode(n *html.Node, words, images *int) {

	if n.Type == html.TextNode {
		*words += len(strings.Fields(n.Data))

	} else if n.Type == html.ElementNode && n.Data == "img" {
		*images += 1
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visitNode(c, words, images)
	}
}

func countWordsAndImages(doc *html.Node) (int, int) {
	var words, images int
	visitNode(doc, &words, &images)
	return words, images
}

// structs
type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func main() {

	root, err := html.Parse(bytes.NewReader([]byte(rawHtml)))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Site is not valid HTML: %v", err)
		os.Exit(1)
	}

	words, pics := countWordsAndImages(root)
	fmt.Printf("%d words and %d images\n", words, pics)

	// ex2
	var album = struct {
		title  string
		artist string
		year   int
		copies int
	}{
		"The White Album",
		"The Beatles",
		1968,
		1000,
	}
}
