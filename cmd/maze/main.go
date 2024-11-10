package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/errors"
)

func main() {
	var width, height int

	var generatorType, solverType string

	flag.IntVar(&width, "width", 30, "Maze width")
	flag.IntVar(&height, "height", 20, "Maze height")
	flag.StringVar(&generatorType, "generatorType", "dfs", "Algorithm for maze generation (dfs/kruskal)")
	flag.StringVar(&solverType, "solverType", "dfs", "Algorithm to solve the maze (astar/prim)")

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println(errors.NewErrNoCommand().Error())
		os.Exit(1)
	}

	application.Init(config.NewConfig(width, height, flag.Args()[len(flag.Args())-1], generatorType, solverType))
}
