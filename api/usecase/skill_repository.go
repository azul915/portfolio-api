package usecase

import (
	"portfolio-api/api/domain/skill"
)

// SkillRepository は、database層のSkillRepositoryのInterface
type SkillRepository interface {
	GetByTerm(string) (skill.Skills, error)
	GetAll() (skill.Skills, error)
	Store(skill.ReqSkill) error
	Delete(skill.DelSkill) error
}
