package usecase

import "github.com/azul915/portfolio-api/src/domain"

type ProductRepository interface {
	FindAll() (domain.Products, error)
}
