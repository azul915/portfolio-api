package usecase

import (
	"log"

	"github.com/azul915/portfolio-api/src/domain"
)

// ProductInteractor は、ProductRepositoryを注入したInteractor
type ProductInteractor struct {
	ProductRepository ProductRepository
}

// Products は、database層のProductRepositoryが集める全ての「products」コレクションを呼び出す
func (interactor *ProductInteractor) Products() (products domain.Products, err error) {

	products, err = interactor.ProductRepository.FindAll()

	if err != nil {
		log.Println(err)
	}

	return

}
