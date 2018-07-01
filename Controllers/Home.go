package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func RegisterRoutes(r *gin.Engine) {
	r.GET("/", Index)
	r.GET("/termites", Termites)
	r.GET("/particles", Particles)
	r.GET("/markovchains", MarkovChains)
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H {
		"title" : "Welcome",
	})
}

func Termites(c *gin.Context) {
	c.HTML(http.StatusOK, "termites.tmpl", gin.H {
		"title" : "Termites",
	})
}

func Particles(c *gin.Context) {
	c.HTML(http.StatusOK, "particles.tmpl", gin.H {
		"title" : "Particles",
	})
}

func MarkovChains(c *gin.Context) {
	c.HTML(http.StatusOK, "markovChains.tmpl", gin.H {
		"title" : "Markov Chains",
	})
}