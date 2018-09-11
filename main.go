package main

import (
	"path/filepath"

	"github.com/james-vaughn/PersonalWebsite/Middleware"
	"github.com/james-vaughn/PersonalWebsite/Services"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/james-vaughn/PersonalWebsite/Controllers"
	"github.com/james-vaughn/PersonalWebsite/Repositories"
)

var (
	repositoryAggregate     *Repositories.Aggregate
	statsService            *Services.StatsService
	pagesService            *Services.PagesService
	homeController          *Controllers.HomeController
	generativeArtController *Controllers.GenerativeArtController
	steganographyController *Controllers.SteganographyController
)

func main() {
	db, err := gorm.Open("sqlite3", DB_NAME)
	if err != nil {
		panic("Failed to connect to Database")
	}
	defer db.Close()

	injectDependencies(db)

	r := createRouter()
	r.Run(PORT)
}

func createRouter() *gin.Engine {
	r := gin.Default()
	addMiddleware(r, statsService)
	setUpRouter(r, homeController, generativeArtController,
		steganographyController)

	return r
}

func addMiddleware(r *gin.Engine, statsService *Services.StatsService) {
	r.Use(gin.Recovery())
	r.Use(Middleware.StatsTracker(statsService))
}

func setUpRouter(r *gin.Engine, controllers ...Controllers.Controller) *gin.Engine {

	r.LoadHTMLGlob("Views/*/*")
	r.HTMLRender = loadTemplates("./Views")
	r.Static("/Public", "./Public")

	for _, controller := range controllers {
		controller.RegisterRoutes(r)
	}

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

func injectDependencies(db *gorm.DB) {
	repositoryAggregate := Repositories.NewAggregate(db)

	statsService = Services.NewStatsService(repositoryAggregate.StatsRepository)
	pagesService = Services.NewPagesService(repositoryAggregate.PagesRepository)

	homeController = Controllers.NewHomeController()
	generativeArtController = Controllers.NewGenerativeArtController(pagesService)
	steganographyController = Controllers.NewStegonographyController(pagesService)
}
