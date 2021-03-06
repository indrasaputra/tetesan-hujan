package usecase

import (
	"context"
	"fmt"
	"strings"

	"github.com/indrasaputra/tetesan-hujan/entity"
	"github.com/pkg/errors"
)

// CreateBookmark defines the business logic to create a new bookmark.
type CreateBookmark interface {
	// Create creates a new bookmark.
	Create(ctx context.Context, bookmark *entity.Bookmark) error
}

// RaindropRepository defines the business logic to save a bookmark in raindrop.io.
type RaindropRepository interface {
	// GetCollections gets all root collections.
	GetCollections(ctx context.Context) ([]*entity.Collection, error)
	// SaveRaindrop saves a bookmark to specific collection in raindrop.io.
	SaveRaindrop(ctx context.Context, bookmark *entity.Bookmark, collectionID int64) error
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

// Create creates a new bookmark.
// First, it will check whether the collection in which the bookmark will be saved exists.
// If it exists, bookmark will be saved. Otherwise, it returns error.
func (rc *RaindropCreator) Create(ctx context.Context, bookmark *entity.Bookmark) error {
	if bookmark == nil {
		return errors.New("Raindrop is nil")
	}

	colls, err := rc.repo.GetCollections(ctx)
	if err != nil {
		return errors.Wrap(err, "GetCollections returns error")
	}

	collID := int64(0)
	for _, coll := range colls {
		if strings.ToLower(coll.Name) == strings.ToLower(bookmark.CollectionName) {
			collID = coll.ID
			break
		}
	}
	if collID == int64(0) {
		return fmt.Errorf("Collection %s is not found", bookmark.CollectionName)
	}

	return rc.repo.SaveRaindrop(ctx, bookmark, collID)
}
