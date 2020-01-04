package usecase

import (
	"log"

	"github.com/azul915/portfolio-api/src/domain"
)

type SkillInteractor struct {
	SkillRepository SkillRepository
}

func (interactor *SkillInteractor) Skills(term string) (skills domain.Skills, err error) {

	skills, err = interactor.SkillRepository.FindAll(term)

	if err != nil {
		log.Println(err)
	}

	return

}
