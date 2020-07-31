package domain

import "time"

// Products は、Productのarray
type Products []Product

// Product は、Name, DemoURL, Feature, Effort, GithubURL, CreatedAt を持つ struct
type Product struct {
	Name      string    `json:"prd_name"`
	DemoURL   string    `json:"demo"`
	Feature   string    `json:"feature"`
	Effort    string    `json:"effort"`
	GithubURL string    `json:"githuburl"`
	CreatedAt time.Time `json:"created_at"`
}

// AddProduct は、CreatedAt, Name, DemoURL, Feature, Effort, GithubURL を持つ struct
type AddProduct struct {
	CreatedAt time.Time `firestore:"created_at" binding:"required"`
	DemoURL   string    `firestore:"demo" binding:"required"`
	Effort    string    `firestore:"effort" binding:"required"`
	Feature   string    `firestore:"feature" binding:"required"`
	GithubURL string    `firestore:"githuburl" binding:"required"`
	Name      string    `firestore:"name" binding:"required"`
}

// DelProduct は、 Name を持つ struct
type DelProduct struct {
	Name string `firestore:"name" binding:"required"`
}
