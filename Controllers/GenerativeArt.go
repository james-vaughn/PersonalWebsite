package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/james-vaughn/PersonalWebsite/Models"
	"github.com/james-vaughn/PersonalWebsite/Services"
)

type GenerativeArtController struct {
	PagesService *Services.PagesService
	pages        []Models.Page
}

const GenerativeArtControllerName = "art"

func NewGenerativeArtController(pagesService *Services.PagesService) *GenerativeArtController {
	return &GenerativeArtController{
		PagesService: pagesService,
		pages:        pagesService.GetPagesFor(GenerativeArtControllerName),
	}
}

func (c *GenerativeArtController) RegisterRoutes(r *gin.Engine) {
	r.GET("/"+GenerativeArtControllerName+"/", c.Index)

	for _, page := range c.pages {
		url := c.PagesService.GetUrlFor(page)
		r.GET(url, c.ProjectPage(page))
	}
}

func (c *GenerativeArtController) Index(context *gin.Context) {

	context.HTML(http.StatusOK, "projectstable.tmpl", gin.H{
		"title":    "Generative Art",
		"projects": c.pages,
	})
}

func (c *GenerativeArtController) ProjectPage(page Models.Page) func(*gin.Context) {
	return func(context *gin.Context) {
		prevPage := c.PagesService.GetPageFromList(c.pages, page.PrevId)
		nextPage := c.PagesService.GetPageFromList(c.pages, page.NextId)

		context.HTML(http.StatusOK, page.Url+".tmpl", gin.H{
			"title":    page.Title,
			"prev_url": c.PagesService.GetUrlFor(prevPage),
			"prev":     prevPage.Title,
			"next_url": c.PagesService.GetUrlFor(nextPage),
			"next":     nextPage.Title,
		})
	}
}
