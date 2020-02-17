package domain

import "time"

// Products は、Productのarray
type Products []Product

// Product は、Name, DemoUrl, Feature, Effort, GithubUrl, CreatedAt を持つ struct
type Product struct {
	Name      string    `json:"prd_name"`
	DemoURL   string    `json:"demo"`
	Feature   string    `json:"feature"`
	Effort    string    `json:"effort"`
	GithubURL string    `json:"githuburl"`
	CreatedAt time.Time `json:"created_at"`
}

// DelProduct は、 Name を持つ struct
type DelProduct struct {
	Name string `firestore:"name" binding:"required"`
}
