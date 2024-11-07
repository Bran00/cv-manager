package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"

	"github.com/Bran00/cv-manager/generated"
	"github.com/Bran00/cv-manager/model"
)

// CreateBasicInfo is the resolver for the createBasicInfo field.
func (r *mutationResolver) CreateBasicInfo(ctx context.Context, basicInfo model.BasicInfoInput) (*model.BasicInfo, error) {
	panic(fmt.Errorf("not implemented: CreateBasicInfo - createBasicInfo"))
}

// UpdateBasicInfo is the resolver for the updateBasicInfo field.
func (r *mutationResolver) UpdateBasicInfo(ctx context.Context, id string, basicInfo *model.BasicInfoInput) (*model.BasicInfo, error) {
	panic(fmt.Errorf("not implemented: UpdateBasicInfo - updateBasicInfo"))
}

// DeleteBasicInfo is the resolver for the deleteBasicInfo field.
func (r *mutationResolver) DeleteBasicInfo(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteBasicInfo - deleteBasicInfo"))
}

// BasicInfo is the resolver for the basicInfo field.
func (r *queryResolver) BasicInfo(ctx context.Context, id string) (*model.BasicInfo, error) {
	panic(fmt.Errorf("not implemented: BasicInfo - basicInfo"))
}

// BasicInfos is the resolver for the basicInfos field.
func (r *queryResolver) BasicInfos(ctx context.Context) ([]*model.BasicInfo, error) {
	panic(fmt.Errorf("not implemented: BasicInfos - basicInfos"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
