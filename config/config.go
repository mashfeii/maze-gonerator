package config

import "github.com/es-debug/backend-academy-2024-go-template/internal/domain/generators"

type Config struct {
	Width         int
	Height        int
	GeneratorType string
	SolverType    string
	Draw          bool
}

func NewConfig(width, height int, generatorType, solverType string, draw bool) *Config {
	return &Config{
		Width:         width,
		Height:        height,
		GeneratorType: generatorType,
		SolverType:    solverType,
		Draw:          draw,
	}
}

func GetGeneratorTypes() map[string]generators.Generator {
	return map[string]generators.Generator{
		"dfs":     generators.DFSGenerator{},
		"kruskal": generators.KruskalGenerator{},
	}
}
