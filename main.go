package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "time"

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

func getSkills(w http.ResponseWriter, r *http.Request) {

    term := "infrastructure"
    skills, err := getSkill(term)
    if err != nil {
        log.Fatalln(err)
    }

    json.NewEncoder(w).Encode(skills)
}

func addSkill(w http.ResponseWriter, r *http.Request) {

    s := Skill{
        Category: Category{
            ID:   2,
            Name: "フレームワーク",
        },
        CreatedAt: time.Now(),
        Detail:    "軽量なAPIを実装したことがある。",
        Duration:  2,
        Name:      "Echo",
        SelfEval:  1,
        Term:      "serverside",
    }
    err := setSkill(s)
    if err != nil {
        log.Fatalln(err)
    }

    json.NewEncoder(w).Encode(err)
}

func deleteSkill(w http.ResponseWriter, r *http.Request) {

    req := delSkill{
        Name: "PHP",
        Term: "serverside",
    }

    err := delete(req)
    if err != nil {
        log.Fatalln(err)
    }

    json.NewEncoder(w).Encode(err)
}

func handleRequests() {
    http.HandleFunc("/", getSkills)
    http.HandleFunc("/add", addSkill)
    http.HandleFunc("/delete", deleteSkill)
    log.Fatal(http.ListenAndServe(":1999", nil))
}

func main() {

    handleRequests()

}