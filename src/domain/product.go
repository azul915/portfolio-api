package domain

import "time"

type Products []Product

type Product struct {
	Name      string    `json:"prd_name"`
	DemoUrl   string    `json:"demo"`
	Feature   string    `json:"feature"`
	Effort    string    `json:"effort"`
	GithubUrl string    `json:"githuburl"`
	CreatedAt time.Time `json:"created_at"`
}

type DelProduct struct {
	Name string `firestore:"name" binding:"required"`
}
