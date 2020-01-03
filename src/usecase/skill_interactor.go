package usecase

import (
	"github.com/azul915/portfolio-api/src/domain"
)

type SkillInteractor struct {
	SkillRepository SkillRepository
}

func (interactor *SkillInteractor) Skills() (skills domain.Skills, err error) {
	skills, err = interactor.SkillRepository.FindAll()
	return
}
