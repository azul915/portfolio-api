package main

import (

    "context"
    "fmt"
    "log"
    "net/http"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"

    "cloud.google.com/go/firestore"
    "firebase.google.com/go"
    "google.golang.org/api/option"
)


func firebaseInit(ctx context.Context) (*firestore.Client, error) {

    // Use a service account
    sa := option.WithCredentialsFile("portfolio-firebase-adminsdk.json")

    // Firebase
    app, err := firebase.NewApp(ctx, nil, sa)
    if err != nil {
        log.Fatalln(err)
        return nil, err
    }

    // Firestoreのclient作成
    client, err := app.Firestore(ctx)
    if err != nil {
        log.Fatalln(err)
        return nil, err
    }

    return client, nil

}


func addSkill(c echo.Context) error {

    err := setServerSideSkill()
    if err != nil {
        log.Fatalln(err)
    }

    return c.JSON(http.StatusOK, err)

}

func getSkills(c echo.Context) error {

    skills, err := getServerSideSkill()
    if err != nil {
        log.Fatalln(err)
    }

    return c.JSON(http.StatusOK, skills)

}

func main() {

    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", getSkills)
    e.GET("/set", addSkill)
    e.Logger.Fatal(e.Start(":1999"))

}


func p(a ...interface{}) {
    fmt.Println(a...)
}