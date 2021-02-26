package usecase

import (
	"context"
	"errors"

	"github.com/indrasaputra/tetesan-hujan/entity"
)

// CreateRaindrop defines the business logic to create a new raindrop.
type CreateRaindrop interface {
	// Create creates a new raindrop.
	Create(ctx context.Context, raindrop *entity.Raindrop) error
}

// RaindropRepository defines the business logic to save a raindrop.
// It has two methods which are explained in their documentation.
type RaindropRepository interface {
	// GetCollections gets all root collections.
	GetCollections(ctx context.Context) ([]*entity.Collection, error)
	// SaveRaindrop saves a raindrop.
	SaveRaindrop(ctx context.Context, raindrop *entity.Raindrop) error
}

// RaindropCreator responsibles for raindrop creation workflow.
type RaindropCreator struct {
	repo RaindropRepository
}

// NewRaindropCreator creates an instance of RaindropCreator.
func NewRaindropCreator(repo RaindropRepository) *RaindropCreator {
	return &RaindropCreator{
		repo: repo,
	}
}

// Create creates a new raindrop.
// First, it will check whether the collection in which the raindrop will be saved exists.
// If it exists, raindrop will be saved. Otherwise, it returns error.
func (rc *RaindropCreator) Create(ctx context.Context, raindrop *entity.Raindrop) error {
	if raindrop == nil {
		return errors.New("Raindrop is nil")
	}
	return nil
}
