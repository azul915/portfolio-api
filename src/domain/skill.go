package domain

import "time"

type Skills []Skill

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

type DelSkill struct {
	Name string `firestore:"name" binding:"required"`
	Term string `firestore:"term" binding:"required"`
}

type AddSkill struct {
	CreatedAt time.Time `firestore:"created_at" binding:"required"`
	Detail    string    `firestore:"detail" binding:"required"`
	Duration  int64     `firestore:"duration" binding:"required"`
	Name      string    `firestore:"name" binding:"required"`
	SelfEval  int64     `firestore:"self_evaluation" binding:"required"`
	Term      string    `firestore:"term" binding:"required"`
}