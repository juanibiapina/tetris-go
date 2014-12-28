package game

type Point struct {
	X int
	Y int
}

type Form []Point

type Block struct {
	rotation int
	forms    []Form
}

func (b *Block) CurrentForm() Form {
	return b.forms[b.rotation]
}

func (b *Block) Rotate() {
	b.rotation = (b.rotation + 1) % len(b.forms)
}

func (b *Block) Unrotate() {
	b.rotation = (b.rotation - 1) % len(b.forms)
	if b.rotation < 0 {
		b.rotation = -b.rotation
	}
}

var AllBlocks []Block = []Block{
	O,
	I,
	T,
	L1,
	L2,
	S1,
	S2,
}

var O = Block{
	forms: []Form{Form{{0, 0}, {1, 0}, {0, 1}, {1, 1}}},
}

var I = Block{
	forms: []Form{
		Form{{-1, 0}, {0, 0}, {1, 0}, {2, 0}},
		Form{{0, -1}, {0, 0}, {0, 1}, {0, 2}},
	},
}

var T = Block{
	forms: []Form{
		Form{{-1, 0}, {0, 0}, {1, 0}, {0, -1}},
		Form{{0, 0}, {0, -1}, {0, 1}, {-1, 0}},
		Form{{-1, 0}, {0, 0}, {1, 0}, {0, 1}},
		Form{{0, 0}, {0, -1}, {0, 1}, {1, 0}},
	},
}

var L1 = Block{
	forms: []Form{
		Form{{0, 0}, {0, -1}, {0, 1}, {1, 1}},
		Form{{-1, 0}, {0, 0}, {1, 0}, {1, -1}},
		Form{{0, 0}, {0, 1}, {0, -1}, {-1, -1}},
		Form{{0, 0}, {-1, 0}, {1, 0}, {-1, 1}},
	},
}

var L2 = Block{
	forms: []Form{
		Form{{0, 0}, {0, -1}, {0, 1}, {-1, 1}},
		Form{{0, 0}, {-1, 0}, {1, 0}, {1, 1}},
		Form{{0, 0}, {0, 1}, {0, -1}, {1, -1}},
		Form{{0, 0}, {-1, 0}, {1, 0}, {-1, -1}},
	},
}

var S1 = Block{
	forms: []Form{
		Form{{0, 0}, {1, 0}, {1, -1}, {0, 1}},
		Form{{0, 0}, {1, 0}, {0, -1}, {-1, -1}},
	},
}

var S2 = Block{
	forms: []Form{
		Form{{0, 0}, {1, 0}, {1, 1}, {0, -1}},
		Form{{0, 0}, {-1, 0}, {0, -1}, {1, -1}},
	},
}
