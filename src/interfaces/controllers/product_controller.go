package controllers

import (
	"log"
	"net/http"

	"github.com/azul915/portfolio-api/src/interfaces/database"
	"github.com/azul915/portfolio-api/src/usecase"
)

type ProductController struct {
	Interactor usecase.ProductInteractor
}

func NewProductController() *ProductController {

	return &ProductController{
		Interactor: usecase.ProductInteractor{
			ProductRepository: &database.ProductRepository{},
		},
	}
}

func (controller *ProductController) Index(c Context) {

	products, err := controller.Interactor.Products()
	if err != nil {
		log.Fatalln(err)
		c.JSON(500, NewError(err))
		return
	}

	c.JSON(http.StatusOK, products)
}
