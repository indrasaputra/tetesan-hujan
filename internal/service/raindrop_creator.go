package service

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
	// ParseURL parses the bookmark's URL to get detailed information of the URL.
	ParseURL(ctx context.Context, url string) (*entity.ParsedURL, error)
	// SaveRaindrop saves a raindrop bookmark to specific collection in raindrop.io.
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

// Create creates a new bookmark.
// First, it will check whether the collection in which the bookmark will be saved exists.
// If it exists, bookmark will be saved. Otherwise, it returns error.
func (rc *RaindropCreator) Create(ctx context.Context, bookmark *entity.Bookmark) error {
	if bookmark == nil {
		return errors.New("Bookmark is nil")
	}

	collID, cerr := rc.getCollectionID(ctx, bookmark.CollectionName)
	if cerr != nil {
		return cerr
	}

	url, perr := rc.parseURL(ctx, bookmark.URL)
	if perr != nil {
		return perr
	}

	rd := createRaindrop(url, bookmark, collID)
	return rc.repo.SaveRaindrop(ctx, rd)
}

func (rc *RaindropCreator) getCollectionID(ctx context.Context, collectionName string) (int64, error) {
	colls, err := rc.repo.GetCollections(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "[GetCollections] returns error")
	}

	collID := int64(0)
	for _, coll := range colls {
		if strings.EqualFold(coll.Name, collectionName) {
			collID = coll.ID
			break
		}
	}
	if collID == int64(0) {
		return 0, fmt.Errorf("collection %s is not found", collectionName)
	}
	return collID, nil
}

func (rc *RaindropCreator) parseURL(ctx context.Context, url string) (*entity.ParsedURL, error) {
	parsed, err := rc.repo.ParseURL(ctx, url)
	if err != nil {
		return nil, errors.Wrap(err, "[ParseURL] returns error")
	}
	if parsed.Error != "" {
		return nil, fmt.Errorf("[ParseURL] URL is invalid/problematic thus get error from Raindrop: %s", parsed.Error)
	}
	return parsed, nil
}

func createRaindrop(url *entity.ParsedURL, bookmark *entity.Bookmark, collectionID int64) *entity.Raindrop {
	rd := &entity.Raindrop{
		Title:        url.Item.Title,
		Excerpt:      url.Item.Excerpt,
		Link:         bookmark.URL,
		CollectionID: collectionID,
	}

	if url.Item.Meta.Canonical != "" {
		rd.Link = url.Item.Meta.Canonical
	}
	return rd
}
