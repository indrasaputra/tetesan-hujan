package usecase

import (
	"context"
	"fmt"
	"strings"

	"github.com/indrasaputra/tetesan-hujan/entity"
	"github.com/pkg/errors"
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
	// SaveRaindrop saves a raindrop to specific collection.
	SaveRaindrop(ctx context.Context, raindrop *entity.Raindrop, collectionID int64) error
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

	colls, err := rc.repo.GetCollections(ctx)
	if err != nil {
		return errors.Wrap(err, "GetCollections returns error")
	}

	collID := int64(0)
	for _, coll := range colls {
		if strings.ToLower(coll.Name) == strings.ToLower(raindrop.CollectionName) {
			collID = coll.ID
			break
		}
	}
	if collID == int64(0) {
		return fmt.Errorf("Collection %s is not found", raindrop.CollectionName)
	}

	return rc.repo.SaveRaindrop(ctx, raindrop, collID)
}
