package main

import (
	"flag"
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
)

func main() {
	var width, height int

	var generatorType, solverType string

	var draw bool

	flag.IntVar(&width, "width", 30, "Maze width")
	flag.IntVar(&height, "height", 20, "Maze height")
	flag.StringVar(&generatorType, "generatorType", "dfs", "Algorithm for maze generation (dfs/kruskal)")
	flag.StringVar(&solverType, "solverType", "dfs", "Algorithm to solve the maze (astar/prim)")

	flag.Parse()
	application.Init(config.NewConfig(width, height, os.Args[1], generatorType, solverType, draw))
}
