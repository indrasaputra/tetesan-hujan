package raindrop_test

import (
	"testing"

	"github.com/indrasaputra/tetesan-hujan/internal/raindrop"
	"github.com/stretchr/testify/assert"
)

func TestNewAPI(t *testing.T) {
	t.Run("successfully create an instance of API", func(t *testing.T) {
		api := raindrop.NewAPI("http://localhost:8080", "random-token")
		assert.NotNil(t, api)
	})
}
