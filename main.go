package main

import (
	"github.com/gin-gonic/gin"
	"./Controllers"
)

func main() {
	const PORT = ":80"

	r := gin.Default()

	r.Use(gin.Recovery())
	r.LoadHTMLGlob("Views/*")

	r.GET("/", Controllers.Index)

	r.Run(PORT)
}