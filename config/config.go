package config

import "github.com/es-debug/backend-academy-2024-go-template/internal/domain/generators"

type Config struct {
	Width         int
	Height        int
	Command       string
	GeneratorType string
	SolverType    string
	Draw          bool
}

func NewConfig(width, height int, command, generatorType, solverType string) *Config {
	return &Config{
		Width:         width,
		Height:        height,
		Command:       command,
		GeneratorType: generatorType,
		SolverType:    solverType,
	}
}

func GetGeneratorTypes() map[string]generators.Generator {
	return map[string]generators.Generator{
		"dfs":     generators.DFSGenerator{},
		"kruskal": generators.KruskalGenerator{},
	}
}
