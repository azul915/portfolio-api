package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"google.golang.org/api/option"

	"github.com/azul915/portfolio-api/src/usecase"
	"github.com/azul915/portfolio-api/src/interfaces/database"
	
)

type SkillController struct {
	Interactor usercase.SkillInteractor
}

func NewSkillsController() *SkillController {
	return &SkillController{
		Interactor: usercase.SkillInteractor{
			SkillRepository: &database.SkillRepository{
				map[string]ineterface{}: m,
			}
		}
	}
}

func (controller *SkillController) Index(c Context) {

	skills, err := controller.Interactor.Skills()
	if err != nil {
		c.JSON(
			http.StatusText(http.StatusInternalServerError),
			NewError(err)
		)
		return
	}

	c.JSON(http.StatusOK, skills)

}