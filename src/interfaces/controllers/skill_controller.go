package controllers

import (
	"net/http"

	"github.com/azul915/portfolio-api/src/usecase"
	"github.com/azul915/portfolio-api/src/interfaces/database"
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