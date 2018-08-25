package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/james-vaughn/PersonalWebsite/Repositories"
)

type HomeController struct{
	Aggregate *Repositories.Aggregate
}

func NewHomeController(aggregate *Repositories.Aggregate) *HomeController {
	return &HomeController{Aggregate: aggregate}
}

func (c *HomeController) RegisterRoutes(r *gin.Engine) {
	r.GET("/", c.Index)
}

func (*HomeController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "homeIndex.tmpl", gin.H {
		"title" : "Welcome",
	})
}