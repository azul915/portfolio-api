package database

import (
	"time"

	"cloud.google.com/go/firestore"

	"portfolio-api/api/domain/skill"
)

// SkillRepository は、SkillドメインについてCloudFirestoreとのやり取りを担うRepository
type SkillRepository struct {
	Fc FirestoreClient
}

// GetByTerm は、引数で受け取ったterm(serverside, frontend, infrastructure)のコレクションについて全てを取得する
func (repo *SkillRepository) GetByTerm(t string) (skills skill.Skills, err error) {

	ctx := repo.Fc.ctx
	client := repo.Fc.client

	// コレクションをカテゴリー順/asc, 自己評価/descでソート、取得
	data := client.Collection(t).
		OrderBy("category", firestore.Asc).
		OrderBy("self_evaluation", firestore.Desc).
		Documents(ctx)

	docs, err := data.GetAll()
	if err != nil {
		return
	}

	skills = make(skill.Skills, 0)
	for _, doc := range docs {
		s := new(skill.Skill)
		mapToStruct(doc.Data(), &s)
		skills = append(skills, *s)
	}

	defer client.Close()

	return

}

// GetAll は、termを跨いで全てのコレクションについて取得する
func (repo *SkillRepository) GetAll() (skills skill.Skills, err error) {

	ctx := repo.Fc.ctx
	client := repo.Fc.client

	// TODO: ["serverside", "frontend", "infrastrucure"]で回しながらGetAllとskillsへのappendを行う
	// 「serverside」コレクションをカテゴリー順/asc, 自己評価/descでソート、取得
	serverside := client.Collection("serverside").
		OrderBy("category", firestore.Asc).
		OrderBy("self_evaluation", firestore.Desc).
		Documents(ctx)

	serversideDocs, err := serverside.GetAll()
	if err != nil {
		return
	}

	skills = make(skill.Skills, 0)
	for _, doc := range serversideDocs {
		s := new(skill.Skill)
		mapToStruct(doc.Data(), &s)
		skills = append(skills, *s)
	}

	// 「frontend」コレクションをカテゴリー順/asc, 自己評価/descでソート、取得
	frontend := client.Collection("frontend").
		OrderBy("category", firestore.Asc).
		OrderBy("self_evaluation", firestore.Desc).
		Documents(ctx)

	frontendDocs, err := frontend.GetAll()
	if err != nil {
		return
	}

	for _, doc := range frontendDocs {
		s := new(skill.Skill)
		mapToStruct(doc.Data(), &s)
		skills = append(skills, *s)
	}

	// 「infrastructure」コレクションをカテゴリー順/asc, 自己評価/descでソート、取得
	infrastructure := client.Collection("infrastructure").
		OrderBy("category", firestore.Asc).
		OrderBy("self_evaluation", firestore.Desc).
		Documents(ctx)

	infrastructureDocs, err := infrastructure.GetAll()
	if err != nil {
		return
	}

	for _, doc := range infrastructureDocs {
		s := new(skill.Skill)
		mapToStruct(doc.Data(), &s)
		skills = append(skills, *s)
	}

	defer client.Close()

	return

}

// Store は、引数で受け取ったSkillについて、該当するterm（serverside, frontend, infrastructure）のコレクションに、新たなドキュメントを追加する
func (repo *SkillRepository) Store(rs skill.ReqSkill) (err error) {

	ctx := repo.Fc.ctx
	client := repo.Fc.client

	// skill配下のcategoryはネストしているため、skill.Categoryのタグ情報を読み込ませて、
	// ①categoryを除くプロパティで先にDocumentを作り、②map[string]でマージする
	as := skill.AddSkill{
		CreatedAt: time.Now(),
		Detail:    rs.Detail,
		Duration:  rs.Duration,
		Name:      rs.Name,
		SelfEval:  rs.SelfEval,
		Term:      rs.Term,
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
			"id":   rs.Category.ID,
			"name": rs.Category.Name,
		},
	}, firestore.MergeAll)
	if err != nil {
		return
	}

	defer client.Close()

	return

}

// Delete は、引数で受け取った値を「skill.DelSkill.Term, skill.DelSkill.Name」として、該当するコレクション内のドキュメントを削除する
func (repo *SkillRepository) Delete(d skill.DelSkill) (err error) {

	ctx := repo.Fc.ctx
	client := repo.Fc.client

	_, err = client.Collection(d.Term).Doc(d.Name).Delete(ctx)
	if err != nil {
		return
	}

	defer client.Close()

	return

}
