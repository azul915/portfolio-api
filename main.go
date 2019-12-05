package main

import (

    _ "fmt"
    _"context"
    "net/http"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"

    _ "cloud.google.com/go/firestore"
    _ "firebase.google.com/go"
    _ "google.golang.org/api/option"

)

func main() {

    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/hello/:username", MainPage())

    e.Start(":1999")

}

func MainPage() echo.HandlerFunc {
    return func(c echo.Context) error {
        username := c.Param("username")
        return c.String(http.StatusOK, "Hello World " + username)
    }
}