package usecase

import "portfolio-api/api/domain"

// ProductRepository は、database層のProductRepositoryのInterface
type ProductRepository interface {
	GetAll() (domain.Products, error)
	Store(domain.Product) error
	Delete(domain.DelProduct) error
}
