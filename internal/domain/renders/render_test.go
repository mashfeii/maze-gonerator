package renders_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/renders"
	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	t.Parallel()
	assert.New(t)

	maze := domain.NewMaze(5, 5)
	bigMaze := domain.NewMaze(10, 10)
	sameMaze := domain.NewMaze(10, 10)

	render := renders.DefaultRenderer{}

	assert.NotPanics(t, func() {
		render.Render(&maze)
	})

	assert.Equal(t, 92, len(render.Render(&maze)))
	assert.Equal(t, 477, len(render.Render(&bigMaze)))
	assert.Equal(t, 477, len(render.Render(&sameMaze)))
}
