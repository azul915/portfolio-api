package usecase

import "portfolio-api/api/domain"

// ProductRepository は、database層のProductRepositoryのInterface
type ProductRepository interface {
	GetAll() (domain.Products, error)
	Store(domain.ReqProduct) error
	Delete(domain.DelProduct) error
}
