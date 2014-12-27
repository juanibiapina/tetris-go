package game

import (
	"math/rand"
	"time"
)

type Game struct {
	r      *rand.Rand
	Blocks []Block
	Board  *Board

	CurrentBlock  *Block
	CurrentBlockX int
	CurrentBlockY int
}

func initRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func New() *Game {
	return &Game{
		Board:  NewBoard(10, 16),
		Blocks: AllBlocks,
		r:      initRand(),
	}
}

func (g *Game) HasCurrentBlock() bool {
	return g.CurrentBlock != nil
}

func (g *Game) CanMoveCurrentBlockDown() bool {
	for r, row := range g.CurrentBlock.Data {
		for c, v := range row {
			if v == 0 {
				continue
			}

			x := g.CurrentBlockX + c
			y := g.CurrentBlockY + r + 1
			if y >= 16 {
				return false
			}
			if g.Board.Tiles[y][x] == 1 {
				return false
			}
		}
	}
	return true
}

func (g *Game) MoveCurrentBlockDown() {
	g.CurrentBlockY += 1
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
	for _, line := range g.Board.Tiles {
		for _, v := range line {
			if v > 1 {
				return true
			}
		}
	}
	return false
}

func (g *Game) Update() bool {
	if g.Over() {
		return false
	}

	if g.HasCurrentBlock() {
		if g.CanMoveCurrentBlockDown() {
			g.MoveCurrentBlockDown()
		} else {
			g.MergeCurrentBlock()
		}
	} else {
		g.spawnBlock()
	}

	return true
}
