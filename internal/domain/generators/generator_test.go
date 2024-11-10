package generators_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/generators"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/errors"
	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	t.Parallel()
	assert.New(t)

	var (
		dfsGenerator     generators.Generator = generators.DFSGenerator{}
		kruskalGenerator generators.Generator = generators.KruskalGenerator{}
	)

	assert.NotPanics(t, func() {
		maze1, err1 := kruskalGenerator.Generate(5, 5)
		maze2, err2 := dfsGenerator.Generate(5, 5)
		maze3, err3 := kruskalGenerator.Generate(100, 100)
		maze4, err4 := dfsGenerator.Generate(100, 100)

		assert.Equal(t, nil, err1)
		assert.Equal(t, nil, err2)
		assert.Equal(t, nil, err3)
		assert.Equal(t, nil, err4)
		assert.NotEqual(t, maze1, maze2)
		assert.NotEqual(t, maze2, maze3)
		assert.NotEqual(t, maze3, maze4)
	})

	_, err := generators.DFSGenerator{}.Generate(0, 0)

	assert.Error(t, errors.ErrSmallSize{}, err)
}
