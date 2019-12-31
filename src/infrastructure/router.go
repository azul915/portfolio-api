package infrastructure

import (
    "github.com/gin-gonic/gin"
    //"github.com/azul915/portfolio-api/src/interfaces"
)

var Router *gin.Engine

func init() {
    r := gin.Default()
    // r.GET("/users", func(c *gin.Context) { skillController.Index(c) })
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    Router = r
}