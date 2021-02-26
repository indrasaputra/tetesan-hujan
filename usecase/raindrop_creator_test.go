package usecase_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/indrasaputra/tetesan-hujan/entity"
	mock_usecase "github.com/indrasaputra/tetesan-hujan/test/mock/usecase"
	"github.com/indrasaputra/tetesan-hujan/usecase"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type RaindropCreator_Executor struct {
	usecase *usecase.RaindropCreator
	repo    *mock_usecase.MockRaindropRepository
}

func TestNewRaindropCreator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of RaindropCreator", func(t *testing.T) {
		exec := createRaindropCreatorExecutor(ctrl)
		assert.NotNil(t, exec.usecase)
	})
}

func TestRaindropCreator_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("empty/nil raindrop is prohibited", func(t *testing.T) {
		exec := createRaindropCreatorExecutor(ctrl)

		err := exec.usecase.Create(context.Background(), nil)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Raindrop is nil")
	})

	t.Run("GetCollections returns error", func(t *testing.T) {
		exec := createRaindropCreatorExecutor(ctrl)

		exec.repo.EXPECT().GetCollections(context.Background()).Return(nil, errors.New("repository closed"))

		raindrop := createValidRaindrop()
		err := exec.usecase.Create(context.Background(), raindrop)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "GetCollections returns error")
	})

	t.Run("collections don't exist", func(t *testing.T) {
		exec := createRaindropCreatorExecutor(ctrl)

		colls := []*entity.Collection{
			{ID: 1, Name: "dummy"},
			{ID: 2, Name: "noname"},
		}
		exec.repo.EXPECT().GetCollections(context.Background()).Return(colls, nil)

		raindrop := createValidRaindrop()
		err := exec.usecase.Create(context.Background(), raindrop)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Collection Learning is not found")
	})

	t.Run("raindrop save returns error", func(t *testing.T) {
		exec := createRaindropCreatorExecutor(ctrl)
		raindrop := createValidRaindrop()

		colls := []*entity.Collection{
			{ID: 1, Name: "dummy"},
			{ID: 2, Name: "Learning"},
		}
		exec.repo.EXPECT().GetCollections(context.Background()).Return(colls, nil)
		exec.repo.EXPECT().SaveRaindrop(context.Background(), raindrop, int64(2)).Return(errors.New("repository closed"))

		err := exec.usecase.Create(context.Background(), raindrop)

		assert.NotNil(t, err)
	})
}

func createValidRaindrop() *entity.Raindrop {
	return &entity.Raindrop{
		CollectionName: "Learning",
		Link:           "http://raindrop.io",
	}
}

func createRaindropCreatorExecutor(ctrl *gomock.Controller) *RaindropCreator_Executor {
	r := mock_usecase.NewMockRaindropRepository(ctrl)
	u := usecase.NewRaindropCreator(r)

	return &RaindropCreator_Executor{
		usecase: u,
		repo:    r,
	}
}
