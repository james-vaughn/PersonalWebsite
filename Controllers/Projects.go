package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProjectsController struct{}

func (c *ProjectsController) RegisterRoutes(r *gin.Engine) {
	r.GET("/projects/", c.Index)
	r.GET("/projects/termites", c.Termites)
	r.GET("/projects/particles", c.Particles)
	r.GET("/projects/markovchains", c.MarkovChains)
}

func (*ProjectsController) Index(context *gin.Context) {
	context.HTML(http.StatusOK, "projectsIndex.tmpl", gin.H {
		"title" : "Welcome",
	})
}

func (*ProjectsController) Termites(context *gin.Context) {
	context.HTML(http.StatusOK, "termites.tmpl", gin.H {
		"title" : "Termites",
		"next_url" : "/projects/particles",
		"next" : "Particles",
	})
}

func (*ProjectsController) Particles(context *gin.Context) {
	context.HTML(http.StatusOK, "particles.tmpl", gin.H {
		"title" : "Particles",
		"next_url" : "/projects/markovchains",
		"next" : "Markov Chains",
		"prev_url" : "/projects/termites",
		"prev" : "Termites",
	})
}

func (*ProjectsController) MarkovChains(context *gin.Context) {
	context.HTML(http.StatusOK, "markovChains.tmpl", gin.H {
		"title" : "Markov Chains",
		"prev_url" : "/projects/particles",
		"prev" : "Particles",
	})
}