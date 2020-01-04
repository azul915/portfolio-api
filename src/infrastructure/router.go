package infrastructure

import (
	"github.com/gin-gonic/gin"

	"github.com/azul915/portfolio-api/src/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	r := gin.Default()

	skillController := controllers.NewSkillController()

	r.GET("/skills", func(c *gin.Context) { skillController.Index("serverside", c) })
	
	Router = r
}
