package infrastructure

import (
    "github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
        "message": "pong",
        })
    })
    // r.Run(":1999")
    Router = router
}