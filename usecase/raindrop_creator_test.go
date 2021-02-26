package usecase_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mock_usecase "github.com/indrasaputra/tetesan-hujan/test/mock/usecase"
	"github.com/indrasaputra/tetesan-hujan/usecase"
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
		assert.Contains(t, "Raindrop is nil", err.Error())
	})
}

func createRaindropCreatorExecutor(ctrl *gomock.Controller) *RaindropCreator_Executor {
	r := mock_usecase.NewMockRaindropRepository(ctrl)
	u := usecase.NewRaindropCreator(r)

	return &RaindropCreator_Executor{
		usecase: u,
		repo:    r,
	}
}
