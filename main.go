package main

import (
	"github.com/gin-gonic/gin"
	"github.com/james-vaughn/PersonalWebsite/Controllers"
)

func main() {
	const PORT = ":80"

	r := gin.Default()

	r.Use(gin.Recovery())
	r.LoadHTMLGlob("Views/*")

	Controllers.RegisterRoutes(r)

	r.Run(PORT)
}