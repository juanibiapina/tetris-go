package sdl

import (
	"github.com/juanibiapina/tetris/game"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	BLOCK_WIDTH = 40
)

type Sdl struct {
	Window    *sdl.Window
	Surface   *sdl.Surface
	lastTicks uint32
}

func New() *Sdl {
	return &Sdl{}
}

func (s *Sdl) Init() {
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, (10+2)*BLOCK_WIDTH, (16+2)*BLOCK_WIDTH, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	s.Window = window
	s.Surface = window.GetSurface()
	s.lastTicks = sdl.GetTicks()
}

func (s *Sdl) Destroy() {
	s.Window.Destroy()
}

func (s *Sdl) Ticks() uint32 {
	current := sdl.GetTicks()
	dt := current - s.lastTicks
	s.lastTicks = current
	return dt
}

func (s *Sdl) Render(g *game.Game) {
	lborder := sdl.Rect{0, 0, BLOCK_WIDTH, 18 * BLOCK_WIDTH}
	s.Surface.FillRect(&lborder, 0x000000ff)

	rborder := sdl.Rect{11 * BLOCK_WIDTH, 0, BLOCK_WIDTH, 18 * BLOCK_WIDTH}
	s.Surface.FillRect(&rborder, 0x000000ff)

	tborder := sdl.Rect{0, 0, 12 * BLOCK_WIDTH, BLOCK_WIDTH}
	s.Surface.FillRect(&tborder, 0x000000ff)

	bborder := sdl.Rect{0, 17 * BLOCK_WIDTH, 12 * BLOCK_WIDTH, BLOCK_WIDTH}
	s.Surface.FillRect(&bborder, 0x000000ff)

	for r, line := range g.Board.Tiles {
		for c, v := range line {
			var color uint32
			if v == 0 {
				color = 0
			} else {
				color = 0x0000ffff
			}
			rect := sdl.Rect{int32((c + 1) * BLOCK_WIDTH), int32((r + 1) * BLOCK_WIDTH), BLOCK_WIDTH, BLOCK_WIDTH}
			s.Surface.FillRect(&rect, color)
		}
	}

	if g.HasBlock() {
		for _, p := range g.CurrentBlock.CurrentForm() {
			x := g.CurrentBlockX + p.X
			y := g.CurrentBlockY + p.Y
			if y < 0 {
				continue
			}
			rect := sdl.Rect{int32((x + 1) * BLOCK_WIDTH), int32((y + 1) * BLOCK_WIDTH), BLOCK_WIDTH, BLOCK_WIDTH}
			s.Surface.FillRect(&rect, 0x0000ffff)
		}
	}

	s.Window.UpdateSurface()
}
