package usecase

import (
	"log"

	"github.com/azul915/portfolio-api/src/domain"
)

type ProductInteractor struct {
	ProductRepository ProductRepository
}

func (interactor *ProductInteractor) Products() (products domain.Products, err error) {

	products, err = interactor.ProductRepository.FindAll()

	if err != nil {
		log.Println(err)
	}

	return

}
