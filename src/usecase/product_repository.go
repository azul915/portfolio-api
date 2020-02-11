package usecase

import "github.com/azul915/portfolio-api/src/domain"

// ProductRepository は、database層のProductRepositoryのInterface
type ProductRepository interface {
	GetAll() (domain.Products, error)
}
