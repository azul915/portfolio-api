package database

import (
	"context"

	"cloud.google.com/go/firestore"

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

	// skill配下のcategoryはネストしているため、domain.Categoryのタグ情報を読み込ませて、
	// ①categoryを除くプロパティで先にDocumentを作り、②map[string]でマージする
	as := domain.AddSkill{
		CreatedAt: s.CreatedAt,
		Detail:    s.Detail,
		Duration:  s.Duration,
		Name:      s.Name,
		SelfEval:  s.SelfEval,
		Term:      s.Term,
	}

	doc := client.Collection(as.Term).Doc(as.Name)

	// ①categoryを除くプロパティでDocumentを作る
	_, err = doc.Set(ctx, as)
	if err != nil {
		return
	}

	// ②map[string]でcategoryを同じDocumentにマージする
	_, err = doc.Set(ctx, map[string]interface{}{
		"category": map[string]interface{}{
			"id":   s.Category.ID,
			"name": s.Category.Name,
		},
	}, firestore.MergeAll)
	if err != nil {
		return
	}

	defer client.Close()

	return

}

func (repo *SkillRepository) Delete(d domain.DelSkill) (err error) {

	ctx := context.Background()

	client, err := firebaseInit(ctx)
	if err != nil {
		return
	}

	_, err = client.Collection(d.Term).Doc(d.Name).Delete(ctx)
	if err != nil {
		return
	}

	defer client.Close()

	return

}
