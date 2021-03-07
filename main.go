package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
)

// Parse RSS url
func Parse() []gin.H {
	var s []gin.H
	url := "https://news.yahoo.co.jp/rss/topics/top-picks.xml"
	feed, _ := gofeed.NewParser().ParseURL(url)
	items := feed.Items
	for _, item := range items {
		h := gin.H{}
		h["title"] = item.Title
		h["link"] = item.Link
		s = append(s, h)
	}
	return s
}

func main() {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", Parse()[0])
	})

	router.Run()
}
