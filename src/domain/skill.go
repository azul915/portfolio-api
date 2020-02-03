package domain

import "time"

type Skills []Skill

type Skill struct {
	Category  Category  `firestore:"category"`
	CreatedAt time.Time `firestore:"created_at"`
	Detail    string    `firestore:"detail"`
	Duration  int64     `firestore:"duration"`
	Name      string    `firestore:"name"`
	SelfEval  int64     `firestore:"self_evaluation"`
	Term      string    `firestore:"term"`
}

type Category struct {
	ID   int64  `firestore:"id"`
	Name string `firestore:"name"`
}