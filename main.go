package main

import (
	"rss/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// urls := []string{"", ""}
	url := ""
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", lib.Parse(url))
	})
	router.Run()
}
