package domain

import "time"

type Skills []Skill

type Skill struct {
	Category  Category  `form:"category" json:"category" firestore:"category" binding:"required"`
	CreatedAt time.Time `form:"created_at" json:"created_at" firestore:"created_at" binding:"required"`
	Detail    string    `form:"detail" json:"detail" firestore:"detail" binding:"required"`
	Duration  int64     `form:"duration" json:"duration" firestore:"duration" binding:"required"`
	Name      string    `form:"name" json:"name" firestore:"name" binding:"required"`
	SelfEval  int64     `form:"self_evaluation" json:"self_evaluation" firestore:"self_evaluation" binding:"required"`
	Term      string    `form:"term" json:"term" firestore:"term" firestore:"term" binding:"required"`
}

type Category struct {
	ID   int64  `form: "id" json:"id" firestore:"id" binding:"required"`
	Name string `form: "name" json:"name" firestore:"name" binding:"required"`
}

type DelSkill struct {
	Name string `form:"name" json:"name" firestore:"name" binding:"required"`
	Term string `form:"term" json:"term" firestore:"term" binding:"required"`
}