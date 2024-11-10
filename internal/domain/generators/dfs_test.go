package generators_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/generators"
	"github.com/stretchr/testify/assert"
)

func TestBreakWall(t *testing.T) {
	t.Parallel()
	assert.New(t)

	n := &domain.Cell{X: 0, Y: 0, WallR: 1, WallB: 0}
	a := &domain.Cell{X: 1, Y: 0, WallR: 1, WallB: 1}

	generators.BreakWallTesting(n, a)

	assert.Equal(t, 0, n.WallB)
	assert.Equal(t, 1, a.WallB)
	assert.Equal(t, 0, n.WallR)
	assert.Equal(t, 1, a.WallR)

	n.WallB = 1
	n.WallR = 1
	a.X = 0
	a.Y = 1

	generators.BreakWallTesting(n, a)

	assert.Equal(t, 0, n.WallB)
	assert.Equal(t, 1, a.WallB)
	assert.Equal(t, 1, n.WallR)
	assert.Equal(t, 1, a.WallR)
}
