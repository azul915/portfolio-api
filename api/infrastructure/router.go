package infrastructure

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"portfolio-api/api/domain"
	"portfolio-api/api/interfaces/controllers"
)

// Router は、server.goから呼び出すGinルーター
var Router *gin.Engine

func init() {
	r := gin.Default()

	s := controllers.NewSkillController()
	r.GET("/skills/:term", func(c *gin.Context) { indexSkillsByTerm(s, c) })
	r.GET("/skills", func(c *gin.Context) { indexAllSkills(s, c) })
	r.POST("/skill", func(c *gin.Context) { createSkill(s, c) })
	r.DELETE("/skill", func(c *gin.Context) { deleteSkill(s, c) })

	p := controllers.NewProductController()
	r.GET("/products", func(c *gin.Context) { indexAllProducts(p, c) })
	r.POST("/product", func(c *gin.Context) { createProduct(p, c) })
	r.DELETE("/product", func(c *gin.Context) { deleteProduct(p, c) })

	Router = r

}

func indexSkillsByTerm(skillController *controllers.SkillController, c *gin.Context) {

	c.Header("access-control-allow-origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")

	term := c.Param("term")

	skillController.IndexByTerm(term, c)

}

func indexAllSkills(skillController *controllers.SkillController, c *gin.Context) {

	c.Header("access-control-allow-origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")

	skillController.Index(c)

}

func createSkill(skillController *controllers.SkillController, c *gin.Context) {

	c.Header("access-control-allow-origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST")

	skill := domain.ReqSkill{}
	if err := c.BindJSON(&skill); err != nil {
		log.Println("BadRequest")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "BadRequest",
		})
		return
	}

	skillController.Create(skill, c)

}

func deleteSkill(skillController *controllers.SkillController, c *gin.Context) {

	c.Header("access-control-allow-origin", "*")
	c.Header("Access-Control-Allow-Methods", "DELETE")

	term := c.Query("term")
	name := c.Query("name")
	deleteSkill := domain.DelSkill{Name: name, Term: term}

	skillController.Delete(deleteSkill, c)

}

func indexAllProducts(productController *controllers.ProductController, c *gin.Context) {

	c.Header("access-control-allow-origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")

	productController.Index(c)

}

func createProduct(productController *controllers.ProductController, c *gin.Context) {

	c.Header("access-control-allow-origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST")

	product := domain.ReqProduct{}
	if err := c.BindJSON(&product); err != nil {
		log.Println("BadRequest")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "BadRequest",
		})
		return
	}
	productController.Create(product, c)
}

func deleteProduct(productController *controllers.ProductController, c *gin.Context) {

	c.Header("access-control-allow-origin", "*")
	c.Header("Access-Control-Allow-Methods", "DELETE")

	name := c.Query("name")
	deleteProduct := domain.DelProduct{Name: name}

	productController.Delete(deleteProduct, c)
}
