package main

import (
	"github.com/gin-gonic/gin"
	//"github.com/gin-contrib/static"
	"github.com/james-vaughn/PersonalWebsite/Controllers"
)

func main() {
	const PORT = ":80"

	r := gin.Default()

	r.Use(gin.Recovery())
	//r.Use(static.Serve("/Public", static.LocalFile("/Public", false)))
	r.LoadHTMLGlob("Views/*")
	r.Static("/Public", "./Public")

	Controllers.RegisterRoutes(r)

	r.Run(PORT)
}