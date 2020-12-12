package skill

import (
	"fmt"
)

// Term は、
type Term interface {
	Name() string
}

// term は、
type term struct{}

type serverside struct {
	*term
}

func (serverside) Name() string {
	return "serverside"
}

type frontend struct {
	*term
}

func (frontend) Name() string {
	return "frontend"
}

type infrastructure struct {
	*term
}

func (infrastructure) Name() string {
	return "infrastructure"
}

// NewTerm は、
func NewTerm(terms string) (Term, error) {

	switch terms {
	case "serverside":
		return &serverside{term: &term{}}, nil
	case "frontend":
		return &frontend{term: &term{}}, nil
	case "infrastructure":
		return &infrastructure{term: &term{}}, nil
	default:
		return nil, fmt.Errorf("%s is invalid", terms)
	}

}
