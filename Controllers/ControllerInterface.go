package Controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	RegisterRoutes(*gin.Engine)
}
