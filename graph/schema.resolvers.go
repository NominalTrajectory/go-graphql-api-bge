package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/NominalTrajectory/go-graphql-api-bge/database"
	"github.com/NominalTrajectory/go-graphql-api-bge/graph/generated"
	"github.com/NominalTrajectory/go-graphql-api-bge/graph/model"
)

var db = database.Connect()

func (r *mutationResolver) AddDevice(ctx context.Context, input *model.NewDevice) (*model.Device, error) {
	return db.AddDevice(input), nil
}

func (r *queryResolver) Device(ctx context.Context, id string) (*model.Device, error) {
	return db.FindDeviceById(id), nil
}

func (r *queryResolver) Devices(ctx context.Context) ([]*model.Device, error) {
	return db.GetAllDevices(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
