package infrastructure

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"portfolio-api/api/domain/product"
	"portfolio-api/api/domain/skill"
	"portfolio-api/api/interfaces/controllers"
)

// Router は、server.goから呼び出すGinルーター
var Router *gin.Engine

func init() {
	r := gin.Default()

	r.GET("/skills/:term", func(c *gin.Context) { indexSkillsByTerm(c) })
	r.GET("/skills", func(c *gin.Context) { indexAllSkills(c) })
	r.POST("/skill", func(c *gin.Context) { createSkill(c) })
	r.DELETE("/skill", func(c *gin.Context) { deleteSkill(c) })

	r.GET("/products", func(c *gin.Context) { indexAllProducts(c) })
	r.POST("/product", func(c *gin.Context) { createProduct(c) })
	r.DELETE("/product", func(c *gin.Context) { deleteProduct(c) })

	Router = r

}

func indexSkillsByTerm(c *gin.Context) {

	skillController := controllers.NewSkillController()

	c.Header("access-control-allow-origin", "http://localhost:8083")
	c.Header("Access-Control-Allow-Methods", "GET")

	term := c.Param("term")

	skillController.IndexByTerm(term, c)

}

func indexAllSkills(c *gin.Context) {

	skillController := controllers.NewSkillController()

	c.Header("access-control-allow-origin", "http://localhost:8083")
	c.Header("Access-Control-Allow-Methods", "GET")

	skillController.Index(c)

}

func createSkill(c *gin.Context) {

	skillController := controllers.NewSkillController()

	c.Header("access-control-allow-origin", "http://localhost:8083")
	c.Header("Access-Control-Allow-Methods", "POST")

	skill := skill.ReqSkill{}
	if err := c.BindJSON(&skill); err != nil {
		log.Println("BadRequest")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "BadRequest",
		})
		return
	}

	skillController.Create(skill, c)

}

func deleteSkill(c *gin.Context) {

	skillController := controllers.NewSkillController()

	c.Header("access-control-allow-origin", "http://localhost:8083")
	c.Header("Access-Control-Allow-Methods", "DELETE")

	term := c.Query("term")
	name := c.Query("name")
	deleteSkill := skill.DelSkill{Name: name, Term: term}

	skillController.Delete(deleteSkill, c)

}

func indexAllProducts(c *gin.Context) {

	productController := controllers.NewProductController()

	c.Header("access-control-allow-origin", "http://localhost:8083")
	c.Header("Access-Control-Allow-Methods", "GET")

	productController.Index(c)

}

func createProduct(c *gin.Context) {

	productController := controllers.NewProductController()

	c.Header("access-control-allow-origin", "http://localhost:8083")
	c.Header("Access-Control-Allow-Methods", "POST")

	product := product.ReqProduct{}
	if err := c.BindJSON(&product); err != nil {
		log.Println("BadRequest")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "BadRequest",
		})
		return
	}
	productController.Create(product, c)
}

func deleteProduct(c *gin.Context) {

	productController := controllers.NewProductController()

	c.Header("access-control-allow-origin", "http://localhost:8083")
	c.Header("Access-Control-Allow-Methods", "DELETE")

	name := c.Query("name")
	deleteProduct := product.DelProduct{Name: name}

	productController.Delete(deleteProduct, c)
}
