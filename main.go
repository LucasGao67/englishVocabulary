package main

import (
	"bufio"
	"fmt"
	"github.com/LucasGao67/englishVocabulary/capture"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var baseUrl = "https://www.collinsdictionary.com/dictionary/english/"

const ltxf = `\documentclass[twocolumn]{book}

\title{English book}
\author{lucas gao}
\date{\today}
\usepackage{lipsum}
\usepackage{color}
\usepackage{sectsty}
\sectionfont{\color{cyan} \fontsize{18}{20}\selectfont}

\begin{document}
`

const ltxa = `\end{document}`

func main() {

	//capture.Convert("away")
	//capture.Convert("test")

	//file, _ := os.Open("word1.txt")
	//defer file.Close()
	//
	//scanner := bufio.NewScanner(file)
	//for scanner.Scan() {
	//	word := scanner.Text()
	//	fmt.Println(word)
	//	capture.Convert(word)
	//	time.Sleep(300 * time.Millisecond)
	//	//fmt.Println(scanner.Text())
	//}
	//captureHtml()
	//_ = convert2Latex()
	genBook("")
}

func captureHtml() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Turn on asynchronous requests
		colly.Async(true),
		// Attach a debugger to the collector
		colly.Debugger(&debug.LogDebugger{}),
	)

	// Limit the number of threads started by colly to two
	// when visiting links which domains' matches "*httpbin.*" glob
	if err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*httpbin.*",
		Parallelism: 2,
		Delay:       300 * time.Millisecond,
	}); err != nil {
		panic(err)
	}

	c.OnRequest(func(request *colly.Request) {

	})

	c.OnResponse(func(response *colly.Response) {
		//fmt.Println(string(response.Body))
		url := response.Request.URL
		names := strings.Split(url.Path, "/")
		name := names[len(names)-1]
		f, _ := os.Create("doc/html/" + name + ".html")
		defer f.Close()

		_, _ = f.WriteString(string(response.Body))
	})

	file, _ := os.Open("doc/wordlist.txt")
	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if err := c.Visit(fmt.Sprintf("%s%s", baseUrl, word)); err != nil {
			fmt.Println(err)
		}
		//capture.Convert(word)
		time.Sleep(300 * time.Millisecond)
		//fmt.Println(scanner.Text())
	}

	// Wait until threads are finished
	c.Wait()
}

func convert2Latex() error {
	files, err := ioutil.ReadDir("doc/html")
	if err != nil {
		return err
	}
	for _, fileInfo := range files {
		//fmt.Println(fileInfo)
		capture.Convert(fileInfo.Name())
	}
	return nil
}

/**
2000词 一本
 */
func genBook(bookType string) {
	file, _ := os.Open("doc/wordlist.txt")
	defer func() {
		_ = file.Close()
	}()
	scanner := bufio.NewScanner(file)

	i := 1
	out, err := os.Create(fmt.Sprintf("doc/book/volumn%d.tex", i/2000+1))
	if err != nil {
		fmt.Println(err)
	}
	out.WriteString(ltxf)
	for scanner.Scan() {
		if i%2000 == 0 {
			out.WriteString(ltxa)
			out.Close()
			out, _ = os.Create(fmt.Sprintf("doc/book/volumn%d.tex", i/2000+1))
			out.WriteString(ltxf)
		}
		word := scanner.Text()
		//fmt.Println(word)
		bytes, _ := ioutil.ReadFile("doc/latex/tiny/" + word + ".txt")

		out.WriteString(string(bytes))
		i++

	}
	defer func() {
		out.WriteString(ltxa)
		out.Close()
	}()
}
