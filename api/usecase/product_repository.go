package usecase

import "portfolio-api/api/domain/product"

// ProductRepository は、database層のProductRepositoryのInterface
type ProductRepository interface {
	GetAll() (product.Products, error)
	Store(product.ReqProduct) error
	Delete(product.DelProduct) error
}
