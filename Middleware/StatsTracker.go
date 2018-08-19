package Middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/james-vaughn/PersonalWebsite/Repositories"
	"github.com/james-vaughn/PersonalWebsite/Models"
	"strings"
)

func StatsTracker(repository *Repositories.Aggregate) gin.HandlerFunc{
	return func (context *gin.Context) {
		//Trying to only grab pages and not resources
		//TODO improve
		if !strings.Contains(context.Request.URL.Path, ".") {
			stat := &Models.Stat{
				IpAddress: context.Request.RemoteAddr,
				Page:      context.Request.URL.Path,
			}

			repository.StatsRepository.Create(stat)
		}
	}
}