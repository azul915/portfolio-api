package usecase

import (
	"log"

	"portfolio-api/api/domain/skill"
)

// SkillInteractor は、SkillRepositoryを注入したInteractor
type SkillInteractor struct {
	SkillRepository SkillRepository
}

// SkillsByTerm は、database層のSkillRepositoryがtermに応じて集めるコレクションを呼び出す
func (interactor *SkillInteractor) SkillsByTerm(term string) (skills skill.Skills, err error) {

	tms, err := skill.NewTerm(term)
	if err != nil {
		return
	}

	skills, err = interactor.SkillRepository.GetByTerm(tms.Name())

	if err != nil {
		log.Println(err)
	}

	return

}

// Skills は、database層のSkillRepositoryのGetAllを呼び出す
func (interactor *SkillInteractor) Skills() (skills skill.Skills, err error) {

	skills, err = interactor.SkillRepository.GetAll()

	if err != nil {
		log.Fatalln(err)
	}

	return

}

// Add は、database層のSkillRepositoryのStoreを呼び出す
func (interactor *SkillInteractor) Add(s skill.ReqSkill) (err error) {

	tms, err := skill.NewTerm(s.Term)
	if err != nil {
		log.Println(err)
		return
	}
	s.Term = tms.Name()

	err = interactor.SkillRepository.Store(s)

	if err != nil {
		log.Println(err)
	}

	return

}

// Delete は、database層のSkillRepositoryのDeleteを呼び出す
func (interactor *SkillInteractor) Delete(s skill.DelSkill) (err error) {

	tms, err := skill.NewTerm(s.Term)
	if err != nil {
		log.Println(err)
		return
	}
	s.Term = tms.Name()

	err = interactor.SkillRepository.Delete(s)

	if err != nil {
		log.Println(err)
	}

	return

}
