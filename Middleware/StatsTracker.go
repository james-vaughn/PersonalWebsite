package Middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/james-vaughn/PersonalWebsite/Models"
	"github.com/james-vaughn/PersonalWebsite/Services"
	"strings"
)

func StatsTracker(statsService *Services.StatsService) gin.HandlerFunc{
	return func (context *gin.Context) {
		//Trying to only grab pages and not resources
		//TODO improve
		if !strings.Contains(context.Request.URL.Path, ".") {
			stat := &Models.Stat{
				IpAddress: context.Request.RemoteAddr,
				Page:      context.Request.URL.Path,
			}

			statsService.Create(stat)
		}
	}
}