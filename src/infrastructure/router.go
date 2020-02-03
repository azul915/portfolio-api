package infrastructure

import (
	"github.com/gin-gonic/gin"

	"github.com/azul915/portfolio-api/src/interfaces/controllers"
	"github.com/azul915/portfolio-api/src/domain"
)

var Router *gin.Engine

func init() {
	r := gin.Default()

	r.GET("/skills", func(c *gin.Context) { indexOfTerm(c) })
	r.POST("/skill", func(c *gin.Context) { createSkill(c) })
	Router = r
}

func indexOfTerm(c *gin.Context) {

	skillController := controllers.NewSkillController()
	c.Header("access-control-allow-origin", "*")
	term := c.Query("term")
	skillController.Index(term, c)

}

func createSkill(c *gin.Context) {

	skillController := controllers.NewSkillController()
	c.Header("access-control-allow-origin", "*")
	var skill = domain.Skill{}
	c.Bind(&skill)
	skillController.Create(skill, c)

}