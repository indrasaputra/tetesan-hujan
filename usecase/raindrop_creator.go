package usecase

import (
	"context"

	"github.com/indrasaputra/tetesan-hujan/entity"
)

// CreateRaindrop defines the business logic to create a new raindrop.
type CreateRaindrop interface {
	// Create creates a new raindrop.
	Create(ctx context.Context, raindrop *entity.Raindrop) error
}

// RaindropCreator responsibles for raindrop creation workflow.
type RaindropCreator struct {
}

// NewRaindropCreator creates an instance of RaindropCreator.
func NewRaindropCreator() *RaindropCreator {
	return &RaindropCreator{}
}
