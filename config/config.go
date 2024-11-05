package config

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
