package usecase

import (
	"log"

	"github.com/azul915/portfolio-api/src/domain"
)

type SkillInteractor struct {
	SkillRepository SkillRepository
}

func (interactor *SkillInteractor) SkillsByTerm(term string) (skills domain.Skills, err error) {

	skills, err = interactor.SkillRepository.FindSkillsByTerm(term)

	if err != nil {
		log.Println(err)
	}

	return

}

func (interactor *SkillInteractor) GetAll() (skills domain.Skills, err error) {

	skills, err = interactor.SkillRepository.GetAll()

	if err != nil {
		log.Fatalln(err)
	}

	return

}

func (interactor *SkillInteractor) Add(skill domain.Skill) (err error) {

	err = interactor.SkillRepository.Store(skill)

	if err != nil {
		log.Println(err)
	}

	return

}

func (interactor *SkillInteractor) Delete(skill domain.DelSkill) (err error) {

	err = interactor.SkillRepository.Delete(skill)

	if err != nil {
		log.Println(err)
	}

	return

}
