package capture

import (
	"github.com/LucasGao67/englishVocabulary/book"
	"github.com/LucasGao67/englishVocabulary/book/entity"
	"github.com/antchfx/htmlquery"
	"io/ioutil"
	"net/http"
	"os"
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

func Convert(word string) {
	//parse1(word)
	doc, _ := htmlquery.LoadURL(baseUrl + word)
	//doc, _ := htmlquery.Parse(strings.NewReader(file))
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
