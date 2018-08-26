package main

import (
	"github.com/james-vaughn/PersonalWebsite/Middleware"
	"github.com/james-vaughn/PersonalWebsite/Services"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/james-vaughn/PersonalWebsite/Controllers"
	"github.com/james-vaughn/PersonalWebsite/Repositories"
)

func main() {
	db, err := gorm.Open("sqlite3", DB_NAME)
	if err != nil {
		panic("Failed to connect to Database")
	}
	defer db.Close()

	repositoryAggregate := Repositories.NewAggregate(db)

	statsService := Services.NewStatsService(repositoryAggregate.StatsRepository)
	pagesService := Services.NewPagesService(repositoryAggregate.PagesRepository)

	controllers := []Controllers.Controller{
		Controllers.NewHomeController(),
		Controllers.NewProjectsController(pagesService),
	}

	r := createRouter(controllers...)
	addMiddleware(r, statsService)

	r.Run(PORT)
}

func createRouter(controllers ...Controllers.Controller) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("Views/*/*")
	r.HTMLRender = loadTemplates("./Views")
	r.Static("/Public", "./Public")

	for _, controller := range controllers {
		controller.RegisterRoutes(r)
	}

	return r
}

func addMiddleware(r *gin.Engine, statsService *Services.StatsService) {
	r.Use(gin.Recovery())
	r.Use(Middleware.StatsTracker(statsService))
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