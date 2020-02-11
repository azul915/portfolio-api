package usecase

import (
	"log"

	"github.com/azul915/portfolio-api/src/domain"
)

// SkillInteractor は、SkillRepositoryを注入したInteractor
type SkillInteractor struct {
	SkillRepository SkillRepository
}

// SkillsByTerm は、database層のSkillRepositoryがtermに応じて集めるコレクションを呼び出す
func (interactor *SkillInteractor) SkillsByTerm(term string) (skills domain.Skills, err error) {

	skills, err = interactor.SkillRepository.GetByTerm(term)

	if err != nil {
		log.Println(err)
	}

	return

}

// Skills は、database層のSkillRepositoryのGetAllを呼び出す
func (interactor *SkillInteractor) Skills() (skills domain.Skills, err error) {

	skills, err = interactor.SkillRepository.GetAll()

	if err != nil {
		log.Fatalln(err)
	}

	return

}

// Add は、database層のSkillRepositoryのStoreを呼び出す
func (interactor *SkillInteractor) Add(skill domain.Skill) (err error) {

	err = interactor.SkillRepository.Store(skill)

	if err != nil {
		log.Println(err)
	}

	return

}

// Delete は、database層のSkillRepositoryのDeleteを呼び出す
func (interactor *SkillInteractor) Delete(skill domain.DelSkill) (err error) {

	err = interactor.SkillRepository.Delete(skill)

	if err != nil {
		log.Println(err)
	}

	return

}
