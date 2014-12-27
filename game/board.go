package game

type Board struct {
	Tiles [][]int
}

func NewBoard(width, height int) *Board {
	tiles := make([][]int, height)

	for i := 0; i < height; i++ {
		tiles[i] = make([]int, width)
	}

	return &Board{
		Tiles: tiles,
	}
}
