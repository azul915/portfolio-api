package controllers

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"google.golang.org/api/option"

	"github.com/gin-gonic/gin"
)

type Skill struct {
	Category  Category  `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	Detail    string    `json:"detail"`
	Duration  int64     `json:"duration"`
	Name      string    `json:"name"`
	SelfEval  int64     `json:"self_evaluation"`
	Term      string    `json:"term"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Skills *[]Skill

func Index(c *gin.Context) {

	ctx := context.Background()

	client, err := firebaseInit(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	data := client.Collection("infrastructure").Documents(ctx)

	docs, err := data.GetAll()
	if err != nil {
		log.Fatalln(err)
	}

	skills := make([]*Skill, 0)
	for _, doc := range docs {
		s := new(Skill)
		mapToStruct(doc.Data(), &s)
		skills = append(skills, s)
	}

	defer client.Close()

	c.JSON(200, skills)

}

func firebaseInit(ctx context.Context) (*firestore.Client, error) {

	sa := option.WithCredentialsFile("credentials.json")

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client, nil

}

// mapから構造体に変換を行う
func mapToStruct(m map[string]interface{}, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}
