package infrastructure

import (
	"github.com/azul915/portfolio-api/src/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	r := gin.Default()

	// skillController := controllers.NewSkillController()
	r.GET("/skills", func(c *gin.Context) { controllers.Index(c) })
	Router = r
}
