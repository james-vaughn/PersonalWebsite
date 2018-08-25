package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (c *HomeController) RegisterRoutes(r *gin.Engine) {
	r.GET("/", c.Index)
}

func (*HomeController) Index(context *gin.Context) {
	context.HTML(http.StatusOK, "homeindex.tmpl", gin.H {
		"title" : "Welcome",
	})
}