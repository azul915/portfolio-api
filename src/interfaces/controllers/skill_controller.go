package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/azul915/portfolio-api/src/domain"
	"github.com/azul915/portfolio-api/src/interfaces/database"
	"github.com/azul915/portfolio-api/src/usecase"
)

// SkillController は、usecase.ProductInteractor をDIした struct
type SkillController struct {
	Interactor usecase.SkillInteractor
}

// NewSkillController は、EntityをDIしたUseCaseをDIしたSkillController
func NewSkillController() *SkillController {

	return &SkillController{
		Interactor: usecase.SkillInteractor{
			SkillRepository: &database.SkillRepository{},
		},
	}
}

// IndexByTerm は、usecase.SkillInteractorのSkillsByTermメソッドの呼び出しを行う
func (controller *SkillController) IndexByTerm(term string, c Context) {

	skills, err := controller.Interactor.SkillsByTerm(term)
	if err != nil {
		log.Fatalln(err)
		c.JSON(500, NewError(err))
		return
	}

	c.JSON(http.StatusOK, skills)

}

// Index は、usecase.SkillInteractorのSkillsの呼び出しを行う
func (controller *SkillController) Index(c Context) {

	skills, err := controller.Interactor.Skills()
	if err != nil {
		log.Fatalln(err)
		c.JSON(500, NewError(err))
		return
	}

	c.JSON(http.StatusOK, skills)

}

// Create は、usecase.SkillInteractorのAddの呼び出しを行う
func (controller *SkillController) Create(s domain.Skill, c Context) {

	err := controller.Interactor.Add(s)
	if err != nil {
		log.Fatalln(err)
		c.JSON(500, NewError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})

}

// Delete は、usercase.SkillInteractorのDeleteの呼び出しを行う
func (controller *SkillController) Delete(d domain.DelSkill, c Context) {

	err := controller.Interactor.Delete(d)
	if err != nil {
		log.Fatalln(err)
		c.JSON(500, NewError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})

}
