package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"portfolio-api/api/domain/product"
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
			ProductRepository: &database.ProductRepository{
				Fc: *database.NewFirestoreClient(),
			},
		},
	}
}

// Index は、usecase.ProductInteractorのProductsメソッドの呼び出しを行う
func (controller *ProductController) Index(c Context) {

	products, err := controller.Interactor.Products()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	c.JSON(http.StatusOK, products)
}

// Create は、usecase.ProductInteractorのAddの呼び出しを行う
func (controller *ProductController) Create(p product.ReqProduct, c Context) {

	err := controller.Interactor.Add(p)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})

}

// Delete は、usecase.ProductInteractorのDeleteメソッドの呼び出しを行う
func (controller *ProductController) Delete(d product.DelProduct, c Context) {

	err := controller.Interactor.Delete(d)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})

}
