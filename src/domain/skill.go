package domain

import (
	"time"
)
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