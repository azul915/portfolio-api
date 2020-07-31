package usecase

import (
	"portfolio-api/api/domain"
)

// SkillRepository は、database層のSkillRepositoryのInterface
type SkillRepository interface {
	GetByTerm(string) (domain.Skills, error)
	GetAll() (domain.Skills, error)
	Store(domain.ReqSkill) error
	Delete(domain.DelSkill) error
}
