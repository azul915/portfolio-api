package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"portfolio-api/api/domain"
	"portfolio-api/api/interfaces/database"
	"portfolio-api/api/usecase"
)

// ProductController は、usecase.ProductInteractor をDIした struct
type ProductController struct {
	Interactor usecase.ProductInteractor
}

// NewProductController は、ProductController を返す
func NewProductController() *ProductController {

	return &ProductController{
		Interactor: usecase.ProductInteractor{
			ProductRepository: &database.ProductRepository{},
		},
	}
}

// Index は、usecase.ProductInteractorのProductsメソッドの呼び出しを行う
func (controller *ProductController) Index(c Context) {

	products, err := controller.Interactor.Products()
	if err != nil {
		log.Fatalln(err)
		c.JSON(500, NewError(err))
		return
	}

	c.JSON(http.StatusOK, products)
}

// Delete は、usecase.ProductInteractorのDeleteメソッドの呼び出しを行う
func (controller *ProductController) Delete(d domain.DelProduct, c Context) {

	err := controller.Interactor.Delete(d)
	if err != nil {
		log.Fatalln(err)
		c.JSON(500, NewError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})

}