# Maze gonerator

This is a simple maze generator that uses a depth-first search/kruskal algorithms to create a maze inside console.

## Installation

```bash
git clone https://github.com/mashfeii/maze-gonerator.git && cd maze-gonerator
make build && cd bin
```

## Usage

- `--width`: Width of the maze
- `--height`: Height of the maze
- `--generatorType`: Type of the generator (`dfs`/`kruskal`)
- `draw/solve`: Draw/solve the maze

```bash
./maze --width 10 --height 10 --generatorType dfs draw
```
