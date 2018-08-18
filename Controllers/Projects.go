package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProjectsRegisterRoutes(r *gin.Engine) {
	r.GET("/projects/", ProjectsIndex)
	r.GET("/projects/termites", Termites)
	r.GET("/projects/particles", Particles)
	r.GET("/projects/markovchains", MarkovChains)
}

func ProjectsIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "projectsIndex.tmpl", gin.H {
		"title" : "Welcome",
	})
}

func Termites(c *gin.Context) {
	c.HTML(http.StatusOK, "termites.tmpl", gin.H {
		"title" : "Termites",
		"next_url" : "/projects/particles",
		"next" : "Particles",
	})
}

func Particles(c *gin.Context) {
	c.HTML(http.StatusOK, "particles.tmpl", gin.H {
		"title" : "Particles",
		"next_url" : "/projects/markovchains",
		"next" : "Markov Chains",
		"prev_url" : "/projects/termites",
		"prev" : "Termites",
	})
}

func MarkovChains(c *gin.Context) {
	c.HTML(http.StatusOK, "markovChains.tmpl", gin.H {
		"title" : "Markov Chains",
		"prev_url" : "/projects/particles",
		"prev" : "Particles",
	})
}