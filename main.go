package main

import (
	"github.com/gin-gonic/gin"
	"github.com/james-vaughn/PersonalWebsite/Controllers"
	"github.com/gin-contrib/multitemplate"
	"path/filepath"
)

func main() {
	r := gin.Default()

	r.Use(gin.Recovery())
	r.LoadHTMLGlob("Views/*/*")
	r.HTMLRender = loadTemplates("./Views")
	r.Static("/Public", "./Public")

	Controllers.HomeRegisterRoutes(r)
	Controllers.ProjectsRegisterRoutes(r)

	r.Run(PORT)
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/Layouts/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/Includes/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		//child template must come first, then the base layouts
		files := append([]string{include}, layouts...)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}