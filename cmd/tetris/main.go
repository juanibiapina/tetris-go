package main

import (
	"github.com/juanibiapina/tetris/game"
	"github.com/juanibiapina/tetris/render/sdl"
	"time"
)

func main() {
	g := game.New()

	renderer := sdl.New()

	renderer.Init()

	run := true
	for run {
		run = g.Update()
		renderer.Render(g)
		time.Sleep(1 * time.Second)
	}

	renderer.Destroy()
}
