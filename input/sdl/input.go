package sdl

import (
	"github.com/juanibiapina/tetris/game"
	"github.com/veandco/go-sdl2/sdl"
)

func Process(g *game.Game) {
	var event sdl.Event

	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.KeyDownEvent:
			switch k := t.Keysym.Sym; k {
			case sdl.K_LEFT:
				g.TryMoveCurrentBlockLeft()
			case sdl.K_RIGHT:
				g.TryMoveCurrentBlockRight()
			case sdl.K_ESCAPE:
				g.End()
			}
		}
	}
}
