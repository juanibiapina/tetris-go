package console

import (
	"github.com/juanibiapina/tetris/render"
)

func New() render.Renderer {
	return &Console{}
}
