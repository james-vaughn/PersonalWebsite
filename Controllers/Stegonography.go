package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/james-vaughn/PersonalWebsite/Models"
	"github.com/james-vaughn/PersonalWebsite/Services"
)

type SteganographyController struct {
	PagesService *Services.PagesService
	pages        []Models.Page
}

const SteganographyControllerName = "steganography"

func NewStegonographyController(pagesService *Services.PagesService) *SteganographyController {
	return &SteganographyController{
		PagesService: pagesService,
		pages:        pagesService.GetPagesFor(SteganographyControllerName),
	}
}

func (c *SteganographyController) RegisterRoutes(r *gin.Engine) {
	r.GET("/"+SteganographyControllerName+"/", c.Index)

	for _, page := range c.pages {
		url := c.PagesService.GetUrlFor(page)
		r.GET(url, c.ProjectPage(page))
	}
}

func (c *SteganographyController) Index(context *gin.Context) {

	context.HTML(http.StatusOK, "projectstable.tmpl", gin.H{
		"title":    "Steganography",
		"projects": c.pages,
	})
}

func (c *SteganographyController) ProjectPage(page Models.Page) func(*gin.Context) {
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
