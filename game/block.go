package game

type Block struct {
	Data  [][]int
	Start int
}

var AllBlocks []Block = []Block{
	Square,
	Line,
	LLeft,
	LRight,
	CurlyL,
	CurlyR,
	T,
}

var Square = Block{
	Data:  [][]int{[]int{1, 1}, []int{1, 1}},
	Start: 4,
}

var Line = Block{
	Data:  [][]int{[]int{1, 1, 1, 1}},
	Start: 3,
}

var LLeft = Block{
	Data:  [][]int{[]int{1, 0}, []int{1, 0}, []int{1, 1}},
	Start: 4,
}

var LRight = Block{
	Data:  [][]int{[]int{0, 1}, []int{0, 1}, []int{1, 1}},
	Start: 4,
}

var CurlyL = Block{
	Data:  [][]int{[]int{1, 0}, []int{1, 1}, []int{0, 1}},
	Start: 4,
}

var CurlyR = Block{
	Data:  [][]int{[]int{0, 1}, []int{1, 1}, []int{1, 0}},
	Start: 4,
}

var T = Block{
	Data:  [][]int{[]int{0, 1, 0}, []int{1, 1, 1}},
	Start: 4,
}
