package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	RegisterRoutes(gr *gin.RouterGroup)
	GetPaged(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}
