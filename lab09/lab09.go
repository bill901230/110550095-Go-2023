package main

import (
	"flag"
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	var id string
	var content string
	var time string
	var id_cnt = 0

	max := flag.Int("max", 10, "Max number of comments to show")
	flag.Parse()

	c := colly.NewCollector()
	c.OnHTML("span.push-userid, span.f3.push-content, span.push-ipdatetime", func(e *colly.HTMLElement) {
		if e.Attr("class") == "f3 hl push-userid" {
			id_cnt += 1
			if id_cnt <= *max {
				id = e.Text
				fmt.Printf("%d. 名字:%s，", id_cnt, id)
			}
		} else if e.Attr("class") == "f3 push-content" {
			if id_cnt <= *max {
				content = e.Text
				fmt.Printf("留言%s，", content)
			}
		} else if e.Attr("class") == "push-ipdatetime" {
			if id_cnt <= *max {
				time = e.Text
				fmt.Printf("時間：%s", time)
			}
		}

	})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL.String())
	})

	// fmt.Printf("%d\n", len(id))

	for i := 0; i < *max; i++ {
		// fmt.Printf("%s\n", id[i])
	}

	c.Visit("https://www.ptt.cc/bbs/joke/M.1481217639.A.4DF.html")
}
