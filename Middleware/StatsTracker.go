package Middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/james-vaughn/PersonalWebsite/Repositories"
	"github.com/james-vaughn/PersonalWebsite/Models"
)

func StatsTracker(repository *Repositories.Aggregate) gin.HandlerFunc{
	return func (context *gin.Context) {
		stat := &Models.Stat{
			IpAddress: context.Request.RemoteAddr,
		}

		repository.StatsRepository.Create(stat)
	}
}