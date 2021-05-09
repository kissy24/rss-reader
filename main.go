package main

import (
	"rss/lib"
	"rss/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	model.Init()
	router.LoadHTMLGlob("templates/*.html")

	//Index
	router.GET("/", func(ctx *gin.Context) {
		xml := model.Select(1)
		ctx.HTML(200, "index.html", lib.Parse(xml.Content))
	})

	//Regist
	router.GET("/register", func(ctx *gin.Context) {
		xmls := model.SelectAll()
		ctx.HTML(200, "register.html", gin.H{
			"xmls": xmls,
		})
	})

	//Create
	router.POST("/new", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		content := ctx.PostForm("content")
		model.Insert(name, content)
		ctx.Redirect(302, "/register")
	})

	//Detail
	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		xml := model.Select(id)
		ctx.HTML(200, "detail.html", gin.H{"xml": xml})
	})

	//Update
	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		name := ctx.PostForm("name")
		content := ctx.PostForm("content")
		model.Update(id, name, content)
		ctx.Redirect(302, "/register")
	})

	//Delete
	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		model.Delete(id)
		ctx.Redirect(302, "/register")

	})

	//Setting RSS
	router.POST("/setting/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		xml := model.Select(id)
		ctx.HTML(200, "index.html", lib.Parse(xml.Content))
	})
	router.Run()
}
