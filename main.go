package main

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/multitemplate"
	"github.com/james-vaughn/PersonalWebsite/Controllers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/james-vaughn/PersonalWebsite/Middleware"
	"github.com/james-vaughn/PersonalWebsite/Repositories"
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to Database")
	}
	defer db.Close()

	aggregate := Repositories.NewAggregate(db)

	r := createRouter(aggregate)
	r.Run(PORT)
}

func createRouter(aggregate *Repositories.Aggregate) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(Middleware.StatsTracker(aggregate))
	r.LoadHTMLGlob("Views/*/*")
	r.HTMLRender = loadTemplates("./Views")
	r.Static("/Public", "./Public")

	Controllers.HomeRegisterRoutes(r)
	Controllers.ProjectsRegisterRoutes(r)

	return r
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