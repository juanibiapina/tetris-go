package game

import (
	"math/rand"
	"time"
)

type Game struct {
	speedup bool
	dt      uint32
	speed   uint32
	running bool
	r       *rand.Rand
	Blocks  []Block
	Board   *Board

	CurrentBlock  *Block
	CurrentBlockX int
	CurrentBlockY int
}

func initRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func New() *Game {
	return &Game{
		Board:   NewBoard(10, 16),
		Blocks:  AllBlocks,
		r:       initRand(),
		running: true,
		speed:   300,
	}
}

func (g *Game) HasBlock() bool {
	return g.CurrentBlock != nil
}

func (g *Game) CanMove(offsetx, offsety int) bool {
	for r, row := range g.CurrentBlock.Data {
		for c, v := range row {
			if v == 0 {
				continue
			}

			x := g.CurrentBlockX + c + offsetx
			y := g.CurrentBlockY + r + offsety
			if x < 0 || x >= 10 {
				return false
			}
			if y >= 16 {
				return false
			}
			if g.Board.Tiles[y][x] >= 1 {
				return false
			}
		}
	}
	return true
}

func (g *Game) CanMoveCurrentBlockDown() bool {
	return g.HasBlock() && g.CanMove(0, 1)
}

func (g *Game) CanMoveCurrentBlockLeft() bool {
	return g.HasBlock() && g.CanMove(-1, 0)
}

func (g *Game) CanMoveCurrentBlockRight() bool {
	return g.HasBlock() && g.CanMove(1, 0)
}

func (g *Game) TryMoveCurrentBlockLeft() {
	if g.CanMoveCurrentBlockLeft() {
		g.MoveCurrentBlockLeft()
	}
}

func (g *Game) TryMoveCurrentBlockRight() {
	if g.CanMoveCurrentBlockRight() {
		g.MoveCurrentBlockRight()
	}
}

func (g *Game) End() {
	g.running = false
}

func (g *Game) MoveCurrentBlockDown() {
	g.CurrentBlockY += 1
}

func (g *Game) MoveCurrentBlockLeft() {
	g.CurrentBlockX -= 1
}

func (g *Game) MoveCurrentBlockRight() {
	g.CurrentBlockX += 1
}

func (g *Game) SpeedUp() {
	g.speedup = true
}

func (g *Game) RestoreSpeed() {
	g.speedup = false
}

func (g *Game) MergeCurrentBlock() {
	for r, row := range g.CurrentBlock.Data {
		for c, v := range row {
			x := g.CurrentBlockX + c
			y := g.CurrentBlockY + r
			g.Board.Tiles[y][x] += v
		}
	}
	g.CurrentBlock = nil
}

func (g *Game) spawnBlock() {
	n := g.r.Intn(len(g.Blocks))
	b := g.Blocks[n]
	g.CurrentBlock = &b
	g.CurrentBlockX = b.Start
	g.CurrentBlockY = 0
}

func (g *Game) Over() bool {
	if !g.running {
		return true
	}

	for _, line := range g.Board.Tiles {
		for _, v := range line {
			if v > 1 {
				return true
			}
		}
	}
	return false
}

func (g *Game) getFilledLines() []int {
	filled := make([]int, 0, 5)

	for r, row := range g.Board.Tiles {
		flag := true
		for _, v := range row {
			if v != 1 {
				flag = false
				break
			}
		}
		if flag {
			filled = append(filled, r)
		}
	}

	return filled
}

func (g *Game) removeLine(line int) {
	for r := line; r > 0; r-- {
		row := g.Board.Tiles[r]
		for c, _ := range row {
			g.Board.Tiles[r][c] = g.Board.Tiles[r-1][c]
		}
	}

	for i := 0; i < 10; i++ {
		g.Board.Tiles[0][i] = 0
	}
}

func (g *Game) DestroyFilledLines() {
	filled := g.getFilledLines()
	n := len(filled)

	if n == 0 {
		return
	}

	for _, line := range filled {
		g.removeLine(line)
	}
}

func (g *Game) Update(dt uint32) {
	g.dt += dt

	threashold := g.speed

	if g.speedup {
		threashold = threashold / 10
	}

	if g.dt < threashold {
		return
	}

	g.dt = 0

	if g.HasBlock() {
		if g.CanMoveCurrentBlockDown() {
			g.MoveCurrentBlockDown()
		} else {
			g.MergeCurrentBlock()
		}
	} else {
		g.spawnBlock()
	}

	g.DestroyFilledLines()
}
