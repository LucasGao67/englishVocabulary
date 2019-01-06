package entity

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strings"
)

var baseAttr = &html.Attribute{
	Key: "class",
	Val: "hi rend-b",
}

func (m *Word) SetForms(node *html.Node) *Word {
	expr := fmt.Sprintf("//*[@id='%s__1']/div[3]/span/span[@class='orth']", m.Name)
	var list []string
	for _, n := range htmlquery.Find(node, expr) {
		content := strings.Trim(htmlquery.InnerText(n), " ")
		list = append(list, content)
	}
	m.Forms = list
	return m
}

func (m *Word) SetContent(node *html.Node) *Word {
	expr := fmt.Sprintf("//*[@id='%s__1']/div[3]/div[@class='hom']", m.Name)
	var list []*Content
	for _, n := range htmlquery.Find(node, expr) {
		content := &Content{}

		// type
		if mtype := htmlquery.FindOne(n, "/span[@class='gramGrp']/span"); mtype != nil {
			content.Type = strings.Trim(htmlquery.InnerText(mtype), " ")
		}

		// sense

		// sentence
		var sentences []string
		for _, n1 := range htmlquery.Find(n, "/div[@class='sense']/div[@class='cit type-example']/span[@class='quote']") {
			sentences = append(sentences, strings.Trim(htmlquery.InnerText(n1), " "))
		}
		content.ExampleSentences = sentences

		var slices []*Slice
		for _, n2 := range htmlquery.Find(n, "/div[@class='sense']/*[1]/node()") {
			//if n2.Type == html.DocumentNode
			slice := &Slice{
				Content:  strings.Trim(htmlquery.InnerText(n2), " "),
				Imported: false,
			}
			if n2.Data == "span" && checkAttr(n2.Attr, baseAttr) {
				slice.Imported = true
			}
			slices = append(slices, slice)
		}
		content.Explain = slices
		if len(content.Explain) != 0 {
			list = append(list, content)
		}
	}
	m.Contents = list
	return m
}

func (m *Word) SetExample(node *html.Node) *Word {
	expr := "//*[@class='assets']/div[1]//span[@class='quote']"
	var list []string
	for _, n := range htmlquery.Find(node, expr) {
		str := strings.Trim(htmlquery.InnerText(n), " ")
		list = append(list, str)
	}
	m.ExampleSentences = list
	return m
}

func checkAttr(attrs []html.Attribute, attr *html.Attribute) bool {
	for _, item := range attrs {
		if item.Key == attr.Key && item.Namespace == attr.Namespace && item.Val == attr.Val {
			return true
		}
	}
	return false
}
