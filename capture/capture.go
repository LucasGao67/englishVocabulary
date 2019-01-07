package capture

import (
	"fmt"
	"github.com/LucasGao67/englishVocabulary/book"
	"github.com/LucasGao67/englishVocabulary/book/entity"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
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
	defer func() {

	}()
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {

	}
	return string(data)
}

func CaptureUrl2Html(word string) string {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(response *colly.Response) {

	})
	return ""
}

func Convert(word string) {
	word = strings.Split(word, ".")[0]
	//doc, _ := htmlquery.LoadURL(baseUrl + word)
	str, err := ioutil.ReadFile("doc/html/" + word + ".html")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(word)
	doc, _ := htmlquery.Parse(strings.NewReader(string(str)))
	w := &entity.Word{}
	w.Name = word

	w.SetExample(doc).SetForms(doc).SetContent(doc)

	w.ShowExample = true
	f, _ := os.Create("doc/latex/all/" + word + ".txt")
	defer f.Close()
	book.Generate(f, w)

	w.ShowExample = false
	f1, _ := os.Create("doc/latex/tiny/" + word + ".txt")
	defer f1.Close()
	book.Generate(f1, w)
}
