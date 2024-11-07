package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"github.com/Bran00/cv-manager/graph/model"
)

// Persons is the resolver for the persons field.
func (r *queryResolver) Persons(ctx context.Context) ([]*model.Person, error) {
	person := &model.Person{
		ID: "id-01",
		Name: "Person 1",
		Age: 15,
		IsMale: false,
	}
	persons := []*model.Person{
		person,
	}
	return persons, nil
}


// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
