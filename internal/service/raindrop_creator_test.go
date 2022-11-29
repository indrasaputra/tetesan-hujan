package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/indrasaputra/tetesan-hujan/entity"
	"github.com/indrasaputra/tetesan-hujan/internal/service"
	mock_service "github.com/indrasaputra/tetesan-hujan/test/mock/service"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type RaindropCreatorExecutor struct {
	service *service.RaindropCreator
	repo    *mock_service.MockRaindropRepository
}

func TestNewRaindropCreator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of RaindropCreator", func(t *testing.T) {
		exec := createRaindropCreatorExecutor(ctrl)
		assert.NotNil(t, exec.service)
	})
}

func TestRaindropCreator_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("empty/nil raindrop is prohibited", func(t *testing.T) {
		exec := createRaindropCreatorExecutor(ctrl)

		err := exec.service.Create(context.Background(), nil)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Bookmark is nil")
	})

	t.Run("GetCollections returns error", func(t *testing.T) {
		exec := createRaindropCreatorExecutor(ctrl)

		exec.repo.EXPECT().GetCollections(context.Background()).Return(nil, errors.New("repository closed"))

		bookmark := createValidBookmark()
		err := exec.service.Create(context.Background(), bookmark)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "[GetCollections] returns error")
	})

	t.Run("collections don't exist", func(t *testing.T) {
		exec := createRaindropCreatorExecutor(ctrl)

		colls := []*entity.Collection{
			{ID: 1, Name: "dummy"},
			{ID: 2, Name: "noname"},
		}
		exec.repo.EXPECT().GetCollections(context.Background()).Return(colls, nil)

		bookmark := createValidBookmark()
		err := exec.service.Create(context.Background(), bookmark)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "collection Learning is not found")
	})

	t.Run("ParseURL process returns error", func(t *testing.T) {
		exec := createRaindropCreatorExecutor(ctrl)
		bookmark := createValidBookmark()

		colls := []*entity.Collection{
			{ID: 1, Name: "dummy"},
			{ID: 2, Name: "Learning"},
		}
		exec.repo.EXPECT().GetCollections(context.Background()).Return(colls, nil)
		exec.repo.EXPECT().ParseURL(context.Background(), bookmark.URL).Return(nil, errors.New("parser closed"))

		err := exec.service.Create(context.Background(), bookmark)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "[ParseURL] returns error")
	})

	t.Run("ParseURL body has an error", func(t *testing.T) {
		exec := createRaindropCreatorExecutor(ctrl)
		bookmark := createValidBookmark()

		colls := []*entity.Collection{
			{ID: 1, Name: "dummy"},
			{ID: 2, Name: "Learning"},
		}
		url := createValidParsedURL()
		url.Error = "try_again"

		exec.repo.EXPECT().GetCollections(context.Background()).Return(colls, nil)
		exec.repo.EXPECT().ParseURL(context.Background(), bookmark.URL).Return(url, nil)

		err := exec.service.Create(context.Background(), bookmark)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "[ParseURL] URL is invalid/problematic thus get error from Raindrop: try_again")
	})

	t.Run("raindrop save returns error", func(t *testing.T) {
		exec := createRaindropCreatorExecutor(ctrl)
		bookmark := createValidBookmark()
		bookmark.CollectionName = "dummy"

		colls := []*entity.Collection{
			{ID: 1, Name: "dummy"},
			{ID: 2, Name: "Learning"},
		}
		url := createValidParsedURL()
		rd := createValidRaindrop(url, bookmark, 1)

		exec.repo.EXPECT().GetCollections(context.Background()).Return(colls, nil)
		exec.repo.EXPECT().ParseURL(context.Background(), bookmark.URL).Return(url, nil)
		exec.repo.EXPECT().SaveRaindrop(context.Background(), rd).Return(errors.New("repository closed"))

		err := exec.service.Create(context.Background(), bookmark)

		assert.NotNil(t, err)
	})

	t.Run("successfully save a raindrop with exact name of collection", func(t *testing.T) {
		exec := createRaindropCreatorExecutor(ctrl)
		bookmark := createValidBookmark()

		colls := []*entity.Collection{
			{ID: 1, Name: "dummy"},
			{ID: 2, Name: "Learning"},
		}
		url := createValidParsedURL()
		rd := createValidRaindrop(url, bookmark, 2)

		exec.repo.EXPECT().GetCollections(context.Background()).Return(colls, nil)
		exec.repo.EXPECT().ParseURL(context.Background(), bookmark.URL).Return(url, nil)
		exec.repo.EXPECT().SaveRaindrop(context.Background(), rd).Return(nil)

		err := exec.service.Create(context.Background(), bookmark)

		assert.Nil(t, err)
	})

	t.Run("successfully save a raindrop with same name of collection", func(t *testing.T) {
		exec := createRaindropCreatorExecutor(ctrl)
		bookmark := createValidBookmark()

		colls := []*entity.Collection{
			{ID: 1, Name: "dummy"},
			{ID: 2, Name: "leArniNg"},
		}
		url := createValidParsedURL()
		rd := createValidRaindrop(url, bookmark, 2)

		exec.repo.EXPECT().GetCollections(context.Background()).Return(colls, nil)
		exec.repo.EXPECT().ParseURL(context.Background(), bookmark.URL).Return(url, nil)
		exec.repo.EXPECT().SaveRaindrop(context.Background(), rd).Return(nil)

		err := exec.service.Create(context.Background(), bookmark)

		assert.Nil(t, err)
	})

	t.Run("successfully save a raindrop with link from bookmark", func(t *testing.T) {
		exec := createRaindropCreatorExecutor(ctrl)
		bookmark := createValidBookmark()

		colls := []*entity.Collection{
			{ID: 1, Name: "dummy"},
			{ID: 2, Name: "Learning"},
		}
		url := createValidParsedURL()
		url.Item.Meta.Canonical = ""
		rd := createValidRaindrop(url, bookmark, 2)

		exec.repo.EXPECT().GetCollections(context.Background()).Return(colls, nil)
		exec.repo.EXPECT().ParseURL(context.Background(), bookmark.URL).Return(url, nil)
		exec.repo.EXPECT().SaveRaindrop(context.Background(), rd).Return(nil)

		err := exec.service.Create(context.Background(), bookmark)

		assert.Nil(t, err)
		assert.Equal(t, bookmark.URL, rd.Link)
	})

	t.Run("successfully save a raindrop with link from canonical", func(t *testing.T) {
		exec := createRaindropCreatorExecutor(ctrl)
		bookmark := createValidBookmark()

		colls := []*entity.Collection{
			{ID: 1, Name: "dummy"},
			{ID: 2, Name: "Learning"},
		}
		url := createValidParsedURL()
		rd := createValidRaindrop(url, bookmark, 2)

		exec.repo.EXPECT().GetCollections(context.Background()).Return(colls, nil)
		exec.repo.EXPECT().ParseURL(context.Background(), bookmark.URL).Return(url, nil)
		exec.repo.EXPECT().SaveRaindrop(context.Background(), rd).Return(nil)

		err := exec.service.Create(context.Background(), bookmark)

		assert.Nil(t, err)
		assert.Equal(t, url.Item.Meta.Canonical, rd.Link)
	})
}

func createValidBookmark() *entity.Bookmark {
	return &entity.Bookmark{
		CollectionName: "Learning",
		URL:            "https://raindrop.io",
	}
}

func createValidParsedURL() *entity.ParsedURL {
	url := &entity.ParsedURL{
		Error: "",
	}
	url.Item.Title = "Raindrop.io website"
	url.Item.Excerpt = "raindrop.io is bookmark saver"
	url.Item.Meta.Canonical = "https://raindrop.io/canonical"

	return url
}

func createValidRaindrop(url *entity.ParsedURL, bookmark *entity.Bookmark, collectionID int64) *entity.Raindrop {
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

func createRaindropCreatorExecutor(ctrl *gomock.Controller) *RaindropCreatorExecutor {
	r := mock_service.NewMockRaindropRepository(ctrl)
	u := service.NewRaindropCreator(r)

	return &RaindropCreatorExecutor{
		service: u,
		repo:    r,
	}
}
