package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeRegisterRoutes(r *gin.Engine) {
	r.GET("/", HomeIndex)
}

func HomeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "homeIndex.tmpl", gin.H {
		"title" : "Welcome",
	})
}