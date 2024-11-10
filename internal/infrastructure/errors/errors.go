package errors

import (
	"fmt"

	"github.com/fatih/color"
)

type ErrTerminalSize struct{}

func NewErrTerminalSize() error {
	return ErrTerminalSize{}
}

func (e ErrTerminalSize) Error() string {
	return "Unable to get terminal size"
}

type ErrSmallSize struct {
	Width  int
	Height int
}

func NewErrSmallSize(width, height int) error {
	return ErrSmallSize{Width: width, Height: height}
}

func (e ErrSmallSize) Error() string {
	return fmt.Sprintf("Window size is too small to generate the maze: widtd=%s, height=%s",
		color.RedString("%d", e.Width), color.RedString("%d", e.Height))
}

type ErrInvalidGenerator struct {
	Generator string
}

func NewErrInvalidGenerator(generator string) error {
	return ErrInvalidGenerator{Generator: generator}
}

func (e ErrInvalidGenerator) Error() string {
	return fmt.Sprintf("Invalid generator type - %s, possible values: dfs, kruskal", color.RedString("%s", e.Generator))
}

type ErrInvalidCommand struct {
	Command string
}

func NewErrInvalidCommand(command string) error {
	return ErrInvalidCommand{Command: command}
}

func (e ErrInvalidCommand) Error() string {
	return fmt.Sprintf("Invalid command %s, possible values: draw, solve", color.RedString("%s", e.Command))
}

type ErrNoCommand struct{}

func NewErrNoCommand() error {
	return ErrNoCommand{}
}

func (e ErrNoCommand) Error() string {
	return fmt.Sprintf("No command provided (possible values: %s, %s)", color.MagentaString("draw"), color.MagentaString("solve"))
}
