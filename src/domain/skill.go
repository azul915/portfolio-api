package domain

import "time"

// Skills は、Skillのarray
type Skills []Skill

// Skill は、Category, CreatedAt, Detail, Duraition, Name, SelfEval, Term を持つ struct
type Skill struct {
	Category  Category  `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	Detail    string    `json:"detail"`
	Duration  int64     `json:"duration"`
	Name      string    `json:"name"`
	SelfEval  int64     `json:"self_evaluation"`
	// TODO: Enumで実装
	Term string `json:"term"`
}

// TODO: Enumで実装
// Category は、ID, Name を持つ struct
type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// DelSkill は、NameとTerm を持つ struct
type DelSkill struct {
	Name string `firestore:"name" binding:"required"`
	Term string `firestore:"term" binding:"required"`
}

// AddSkill は、CreatedAt, Detail, Duraition, Name, SelfEval, Term を持つ struct
type AddSkill struct {
	CreatedAt time.Time `firestore:"created_at" binding:"required"`
	Detail    string    `firestore:"detail" binding:"required"`
	Duration  int64     `firestore:"duration" binding:"required"`
	Name      string    `firestore:"name" binding:"required"`
	SelfEval  int64     `firestore:"self_evaluation" binding:"required"`
	// TODO: Enumで実装
	Term string `firestore:"term" binding:"required"`
}
