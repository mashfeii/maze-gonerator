package application_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/errors"
	"github.com/stretchr/testify/assert"
)

func TestValidateSize(t *testing.T) {
	t.Parallel()
	assert.New(t)

	smallConfig := config.NewConfig(2, 2, "draw", "dfs", "astar")

	assert.Error(t, errors.ErrSmallSize{Width: 2, Height: 2}, application.TestingValidateSize(smallConfig))
}
