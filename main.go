package main

import (
    "net/http"
    "fmt"

    "github.com/labstack/echo"

    "golang.org/x/net/context"
    firebase "firebase.google.com/go"

    "google.golang.org/api/option"
)

func initializeAppWithServiceAccount() *firebase.App {

    opt := option.WithCredentialsFile("./portfolio-firebase-adminsdk.json")
    app, err := firebase.NewApp(context.Background(), nil, opt)
    if err != nil {
        log.Fatalf("error initializing app: %v\n", err)
    }

    return app
}

func main() {

    app := initializeAppWithServiceAccount()
    fmt.Print(app)
    if err != nil {
        panic(err)
    }
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "hoge!")
    })
    e.Logger.Fatal(e.Start(":1999"))
}