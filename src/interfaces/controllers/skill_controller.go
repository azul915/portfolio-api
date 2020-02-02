package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/azul915/portfolio-api/src/domain"
	"github.com/azul915/portfolio-api/src/interfaces/database"
	"github.com/azul915/portfolio-api/src/usecase"
)

type SkillController struct {
	Interactor usecase.SkillInteractor
}

func NewSkillController() *SkillController {

	return &SkillController {
		Interactor: usecase.SkillInteractor {
			SkillRepository: &database.SkillRepository {},
		},
	}
}

func (controller *SkillController) Index(term string, c Context) {

	skills, err := controller.Interactor.Skills(term)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	c.JSON(http.StatusOK, skills)

}

func (controller *SkillController) Create(s domain.Skill, c Context) {

	err := controller.Interactor.Add(s)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}