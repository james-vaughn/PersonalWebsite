package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/james-vaughn/PersonalWebsite/Models"
	"github.com/james-vaughn/PersonalWebsite/Services"
	"net/http"
)

type ProjectsController struct{
	PagesService *Services.PagesService
	pages        []Models.Page
}

const ControllerName = "projects"

func NewProjectsController(pagesService *Services.PagesService) *ProjectsController {
	return &ProjectsController{
		PagesService: pagesService,
		pages :       pagesService.GetPagesFor(ControllerName),
	}
}

func (c *ProjectsController) RegisterRoutes(r *gin.Engine) {
	r.GET("/"+ControllerName, c.Index)

	for _, page := range c.pages {
		url := c.PagesService.GetUrlFor(page)
		r.GET(url, c.ProjectPage(page))
	}
}

func (c *ProjectsController) Index(context *gin.Context) {

	context.HTML(http.StatusOK, "projectsindex.tmpl", gin.H {
		"title" : "Projects",
		"projects" : c.pages,
	})
}

func (c *ProjectsController) ProjectPage(page Models.Page) func(*gin.Context){
	return func(context *gin.Context) {
		prevPage := c.PagesService.GetPageFromList(c.pages, page.PrevId)
		nextPage := c.PagesService.GetPageFromList(c.pages, page.NextId)

		context.HTML(http.StatusOK, page.Url+".tmpl", gin.H {
			"title" : page.Title,
			"prev_url" : c.PagesService.GetUrlFor(prevPage),
			"prev" : prevPage.Title,
			"next_url" : c.PagesService.GetUrlFor(nextPage),
			"next" : nextPage.Title,
		})
	}
}