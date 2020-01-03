package usecase

import (
	"github.com/azul915/portfolio-api/src/domain"
)

type SkillRepository interface {
	FindAll() (domain.Skills, error)
}