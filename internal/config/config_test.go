package config_test

import (
	"testing"

	"github.com/indrasaputra/tetesan-hujan/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	t.Run("fail to create an instance of Config due to incomplete env", func(t *testing.T) {
		cfg, err := config.NewConfig("../../test/fixture/env.incomplete")

		assert.NotNil(t, err)
		assert.Nil(t, cfg)
	})

	t.Run("successfully create an instance of Config", func(t *testing.T) {
		cfg, err := config.NewConfig("../../test/fixture/env.valid")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
	})
}
