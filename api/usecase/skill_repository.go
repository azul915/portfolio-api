package usecase

import (
	"portfolio-api/api/domain"
)

// SkillRepository は、database層のSkillRepositoryのInterface
type SkillRepository interface {
	GetByTerm(string) (domain.Skills, error)
	GetAll() (domain.Skills, error)
	Store(domain.Skill) error
	Delete(domain.DelSkill) error
}
