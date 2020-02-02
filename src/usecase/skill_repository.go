package usecase

import (
	"github.com/azul915/portfolio-api/src/domain"
)

type SkillRepository interface {
	FindAll(string) (domain.Skills, error)
	Store(domain.Skill) (error)
}