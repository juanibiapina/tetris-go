package render

import (
	"github.com/juanibiapina/tetris/game"
)

type Renderer interface {
	Init()
	Destroy()
	Render(g *game.Game)
}
