package infrastructure

import (
	"github.com/gin-gonic/gin"
	"time"

	"github.com/azul915/portfolio-api/src/interfaces/controllers"
	"github.com/azul915/portfolio-api/src/domain"
)

var Router *gin.Engine

func init() {
	r := gin.Default()

	r.GET("/skills", func(c *gin.Context) { indexOfTerm(c) })
	r.GET("/skill", func(c *gin.Context) { createSkill(c) })
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

	skill := domain.Skill {
		Category: domain.Category {
			ID: 0,
			Name: "言語",
		},
		CreatedAt: time.Now(),
		Detail: "SpringBootでAPIコンテナを実装したことがある。",
		Duration: 3,
		Name: "Kotlin",
		SelfEval: 4,
		Term: "serverside",
	}

	skillController.Create(skill, c)

}