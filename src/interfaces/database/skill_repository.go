package database

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"google.golang.org/api/option"

	"github.com/azul915/portfolio-api/src/domain"
)

type SkillRepository struct {
	Val interface{}
}

func (repo *SkillRepository) FindAll(t string) (skills domain.Skills, err error) {

	ctx := context.Background()

	client, err := firebaseInit(ctx)
	if err != nil {
		return
	}

	data := client.Collection(t).Documents(ctx)

	docs, err := data.GetAll()
	if err != nil {
		return
	}

	skills = make(domain.Skills, 0)
	for _, doc := range docs {
		s := new(domain.Skill)
		mapToStruct(doc.Data(), &s)
		skills = append(skills, *s)
	}

	defer client.Close()

	return

}

func (repo *SkillRepository) Store(s domain.Skill) (err error) {

	ctx := context.Background()

	client, err := firebaseInit(ctx)
	if err != nil {
		return
	}

	collection := client.Collection(s.Term)

	doc := collection.Doc(s.Name)
	_, err = doc.Set(ctx, s)

	if err != nil {
		return
	}

	defer client.Close()

	return

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