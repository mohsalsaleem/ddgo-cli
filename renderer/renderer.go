package renderer

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mohsalsaleem/ddgo-cli/utils"
	"golang.org/x/net/html"
)

type result struct {
	title string
	url   string
	desc  string
}

func findResultsWrapperNode(doc *html.Node) (*html.Node, error) {
	var linksNode *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode {
			attributes := node.Attr
			for _, attribute := range attributes {
				if attribute.Key == "id" && attribute.Val == "links" {
					linksNode = node
					return
				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if linksNode != nil {
		return linksNode, nil
	}
	return nil, errors.New("No search results found")
}

func getTextFromNode(node *html.Node) string {
	var textContent string = ""
	var addText func(string)
	addText = func(text string) {
		if len(textContent) == 0 {
			textContent = text
		} else {
			textContent = textContent + text
		}
	}
	nodeString := utils.HTMLToString(node)
	tokenizer := html.NewTokenizer(strings.NewReader(nodeString))
loopTokens:
	for {
		token := tokenizer.Next()
		switch token {
		case html.ErrorToken:
			break loopTokens
		case html.TextToken:
			addText(strings.TrimSpace(html.UnescapeString(string(tokenizer.Text()))))
		}
	}
	return textContent
}

func getTextByClassName(node *html.Node, className string) string {
	var text string = ""
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode {
			attributes := node.Attr
			attrMap := make(map[string]string)
			for _, attribute := range attributes {
				attrMap[attribute.Key] = attribute.Val
			}
			if val, ok := attrMap["class"]; ok {
				if strings.Contains(val, className) {
					text = getTextFromNode(node)
					return
				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(node)
	return text
}

func getLinkByClassName(node *html.Node, className string) string {
	var text string = ""
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			attributes := node.Attr
			attrMap := make(map[string]string)
			for _, attribute := range attributes {
				attrMap[attribute.Key] = attribute.Val
			}
			if val, ok := attrMap["class"]; ok {
				if strings.Contains(val, className) {
					text = attrMap["href"]
					return
				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(node)
	return text
}

func isValidResultNode(node *html.Node) bool {
	attributes := node.Attr
	attrMap := make(map[string]string)
	for _, attribute := range attributes {
		attrMap[attribute.Key] = attribute.Val
	}
	if val, ok := attrMap["class"]; ok {
		if strings.Contains(val, "result") {
			return true
		}
	}
	return false
}

func collectResults(linksNode *html.Node) []result {
	results := make([]result, 0)
	for child := linksNode.FirstChild; child != nil; child = child.NextSibling {
		if isValidResultNode(child) {
			var title string = getTextByClassName(child, "result__a")
			var desc string = getTextByClassName(child, "result__snippet")
			var url string = getLinkByClassName(child, "result__url")
			res := result{title: title, url: url, desc: desc}
			results = append(results, res)
		}
	}
	return results
}

func extractResults(htm string) ([]result, error) {
	doc, err := html.Parse(strings.NewReader(htm))
	if err != nil {
		return nil, err
	}

	linksNode, err := findResultsWrapperNode(doc)
	if err != nil {
		return nil, err
	}

	utils.WriteStringToFile(utils.HTMLToString(linksNode), "/tmp/ddgo.html")

	results := collectResults(linksNode)
	return results, nil
}

// Render - Display results
func Render(htm string) {
	results, err := extractResults(htm)
	if err != nil {
		fmt.Println(err.Error())
	}
	for idx, result := range results {
		fmt.Printf("\u001b[37m%d. %s\n", (idx + 1), result.title)
		fmt.Printf("\u001b[34m%s\n", result.url)
		fmt.Printf("\u001b[37m%s\n", result.desc)
		fmt.Println("")
	}
}
