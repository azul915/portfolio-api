package domain

import "time"

type Skills []Skill

type Skill struct {
	Category  Category  `firestore:"category" binding:"required"`
	CreatedAt time.Time `firestore:"created_at" binding:"required"`
	Detail    string    `firestore:"detail" binding:"required"`
	Duration  int64     `firestore:"duration" binding:"required"`
	Name      string    `firestore:"name" binding:"required"`
	SelfEval  int64     `firestore:"self_evaluation" binding:"required"`
	Term      string    `firestore:"term" binding:"required"`
}

type Category struct {
	ID   int64  `firestore:"id" binding:"required"`
	Name string `firestore:"name" binding:"required"`
}

type DelSkill struct {
	Name      string     `firestore:"name" binding:"required"`
	Term      string     `firestore:"term" binding:"required"`
}