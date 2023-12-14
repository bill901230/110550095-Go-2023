package main

import (
	"flag"
	"fmt"

	"github.com/gocolly/colly"
)

var id string
var content string
var time string

func main() {
	cnt := 0
	c := colly.NewCollector()

	max := flag.Int("max", 10, "Max number of comments to show")
	flag.Parse()
	c.OnHTML("span.f3.push-userid, span.f3.push-content, span.push-ipdatetime", func(e *colly.HTMLElement) {
		// fmt.Printf("%s\n", e.Attr("class"))

		if e.Attr("class") == "f3 hl push-userid" {
			cnt += 1
			if cnt <= *max {
				id = e.Text
				fmt.Printf("%d. 名字：%s，", cnt, id)
			}
		} else if e.Attr("class") == "f3 push-content" {
			if cnt <= *max {
				content = e.Text
				fmt.Printf("留言%s，", content)
			}
		} else if e.Attr("class") == "push-ipdatetime" {
			if cnt <= *max {
				time = e.Text
				fmt.Printf("時間：%s", time)
			}
		}
	})

	c.Visit("https://www.ptt.cc/bbs/joke/M.1481217639.A.4DF.html")
}
