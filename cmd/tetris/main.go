package main

import (
	"github.com/juanibiapina/tetris/game"
	input "github.com/juanibiapina/tetris/input/sdl"
	render "github.com/juanibiapina/tetris/render/sdl"
)

func main() {
	g := game.New()

	sdl := render.New()

	sdl.Init()

	var acc uint32 = 0
	for {
		dt := sdl.Ticks()

		input.Process(g)

		g.Update(dt)

		if g.Over() {
			break
		}

		acc += dt

		if acc > 16 {
			acc = 0
			sdl.Render(g)
		}
	}

	sdl.Destroy()
}
