package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct{}

func (c *HomeController) RegisterRoutes(r *gin.Engine) {
	r.GET("/", c.Index)
}

func (*HomeController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "homeIndex.tmpl", gin.H {
		"title" : "Welcome",
	})
}