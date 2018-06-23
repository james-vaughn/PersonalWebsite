package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.Use(gin.Recovery())
	r.LoadHTMLGlob("Views/*")

	r.GET("/", Index)

	r.Run(":8080")
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H {
		"title" : "Welcome",
	})
}