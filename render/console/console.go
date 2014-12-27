package console

import (
	"fmt"
	"github.com/juanibiapina/tetris/game"
)

type Console struct {
}

func (c *Console) Init() {
}

func (c *Console) Destroy() {
}

func (c *Console) Render(g *game.Game) {
	board := game.NewBoard(10, 16)
	for l, _ := range board.Tiles {
		copy(board.Tiles[l], g.Board.Tiles[l])
	}

	if g.HasCurrentBlock() {
		for r, row := range g.CurrentBlock.Data {
			for c, v := range row {
				x := g.CurrentBlockX + c
				y := g.CurrentBlockY + r
				board.Tiles[y][x] += v
			}
		}
	}

	width := len(board.Tiles[0])
	for i := 0; i < width+2; i++ {
		fmt.Print("\u2587")
	}
	fmt.Println()

	for _, line := range board.Tiles {
		fmt.Print("\u2587")

		for _, v := range line {
			if v == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("\u2587")
			}
		}

		fmt.Print("\u2587")
		fmt.Println()
	}

	for i := 0; i < width+2; i++ {
		fmt.Print("\u2587")
	}
	fmt.Println()
}
