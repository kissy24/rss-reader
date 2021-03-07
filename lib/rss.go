package lib

import (
	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
)

// RSSの構造体
type RSS struct {
	Title string
	Link  string
}

// Parse RSS url
func Parse(url string) gin.H {
	feed, _ := gofeed.NewParser().ParseURL(url)
	items := feed.Items
	var rss []RSS
	for _, item := range items {
		rss = append(rss, RSS{item.Title, item.Link})
	}
	return gin.H{"rss": rss}
}
