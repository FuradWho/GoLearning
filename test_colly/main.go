package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

func main() {

	log.Infoln("start colly collector")
	c := colly.NewCollector()

	pageList := make(map[string]string)

	c.OnHTML("dd", func(e *colly.HTMLElement) {
		if e.Request.URL.String() == "http://www.b520.cc/86_86700/" {
			pageList[e.Text] = e.ChildAttr("a[href]", "href")
		}

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("http://www.b520.cc/86_86700/")

	for k, v := range pageList {
		fmt.Printf("page title :%s -> link :%s \n", k, v)

		c.OnHTML("#content", func(e *colly.HTMLElement) {
			var str []byte
			e.DOM.Find("p").Each(func(_ int, selection *goquery.Selection) {
				str = append(str, []byte(fmt.Sprintf("%s\n", selection.Text()))...)
			})
			fileDir := "../novels/" + k + ".txt"
			err := ioutil.WriteFile(fileDir, str, 0666)
			if err != nil {
				return
			}
		})

		c.Visit(v)
	}

}
