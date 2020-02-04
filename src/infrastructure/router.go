package infrastructure

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/azul915/portfolio-api/src/domain"
	"github.com/azul915/portfolio-api/src/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	r := gin.Default()

	r.GET("/skills", func(c *gin.Context) { indexOfTerm(c) })
	r.POST("/skill", func(c *gin.Context) { createSkill(c) })
	r.DELETE("/skill", func(c *gin.Context) { deleteSkill(c) })
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

	if err := c.Bind(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "BadRequest",
		})
		return
	}

	skillController.Create(skill, c)

}

func deleteSkill(c *gin.Context) {

	skillController := controllers.NewSkillController()
	c.Header("access-control-allow-origin", "*")

	term := c.Query("term")
	name := c.Query("name")
	deleteSkill := domain.DelSkill{Name: name, Term: term,}

	skillController.Delete(deleteSkill, c)

}