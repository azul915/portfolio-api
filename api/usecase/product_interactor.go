package usecase

import (
	"log"

	"portfolio-api/api/domain/product"
)

// ProductInteractor は、ProductRepositoryを注入したInteractor
type ProductInteractor struct {
	ProductRepository ProductRepository
}

// Products は、database層のProductRepositoryが集める全ての「products」コレクションを呼び出す
func (interactor *ProductInteractor) Products() (products product.Products, err error) {

	products, err = interactor.ProductRepository.GetAll()

	if err != nil {
		log.Println(err)
	}

	return

}

// Add は、database層のProductRepositoryのStoreを呼び出す
func (interactor *ProductInteractor) Add(product product.ReqProduct) (err error) {

	err = interactor.ProductRepository.Store(product)

	if err != nil {
		log.Println(err)
	}

	return
}

// Delete は、database層のProductRepositoryのDeleteを呼び出す
func (interactor *ProductInteractor) Delete(product product.DelProduct) (err error) {

	err = interactor.ProductRepository.Delete(product)

	if err != nil {
		log.Println(err)
	}

	return

}
