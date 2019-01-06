package capture

import (
	"fmt"
	"github.com/LucasGao67/englishVocabulary/book"
	"github.com/LucasGao67/englishVocabulary/book/entity"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"gopkg.in/xmlpath.v2"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var baseUrl = "https://www.collinsdictionary.com/dictionary/english/"

func captureUrl(word string) string {
	resp, err := http.Get(baseUrl + word)
	if err != nil {

	}
	defer resp.Body.Close()
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {

	}
	return string(data)
}

func parse(html string) *entity.Word {
	path, _ := xmlpath.Compile("//*[@id='challenge__1']/div[3]/span")

	root, err := xmlpath.Parse(strings.NewReader(html))
	if err != nil {
		//log.Fatal(err)
	}
	if value, ok := path.String(root); ok {
		fmt.Println("Found:", value)
	}
	return nil
}

func parse1(word string) {
	doc, err := htmlquery.LoadURL(baseUrl + word)
	if err != nil {
		panic(err)
	}

	for _, n := range htmlquery.Find(doc, "//*[@id='challenge__1']/div[3]/span/span[@class='orth']") {
		//fmt.Println(n)
		fmt.Println(strings.Trim(htmlquery.InnerText(n), " "))
	}
	//htmlquery.Find(n, "")
}

func parseFrom(node *html.Node, word string) [] string {
	expr := fmt.Sprintf("//*[@id='%s__1']/div[3]/span/span[@class='orth']", word)
	var list []string
	for _, n := range htmlquery.Find(node, expr) {
		//fmt.Println(n)
		content := strings.Trim(htmlquery.InnerText(n), " ")
		list = append(list, content)
		fmt.Println(content)
	}
	return list
}

func parseContents(node *html.Node, word string) {

}

func Convert(word string) {
	//parse1(word)
	//doc, _ := htmlquery.LoadURL(baseUrl + word)
	doc, _ := htmlquery.Parse(strings.NewReader(file))
	w := &entity.Word{}
	w.Name = word
	w.SetExample(doc).SetForms(doc).SetContent(doc)

	w.ShowExample = true
	f, _ := os.Create("generate/" + word + ".txt")
	defer f.Close()
	book.Generate(f, w)
}
