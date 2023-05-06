package main

import "fmt"

const (
	FIELD_WIDTH  = 13
	FIELD_HEIGHT = 21
)

type Field [FIELD_HEIGHT][FIELD_WIDTH]int

type Position struct {
	x int
	y int
}

func (p *Position) Init() {
	p.x = 4
	p.y = 0
}

type Game struct {
	field Field
	pos   Position
	block blockKind
}

func NewGame() *Game {
	g := new(Game)

	g.field = [FIELD_HEIGHT][FIELD_WIDTH]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
	g.pos.Init()
	g.block = distribution{}.BlockKind()

	return g
}

func IsCollision(field Field, pos Position, block blockKind) bool {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if y+pos.y >= FIELD_HEIGHT || x+pos.x >= FIELD_WIDTH {
				continue
			}
			if (field[y+pos.y][x+pos.x] & BLOCKS[block][y][x]) == 1 {
				return true
			}
		}
	}
	return false
}

func (g *Game) Draw() {
	fieldBuf := g.field
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if BLOCKS[g.block][y][x] == 1 {
				fieldBuf[y+g.pos.y][x+g.pos.x] = 1
			}
		}
	}

	fmt.Print("\x1b[H")
	for y := 0; y < FIELD_HEIGHT; y++ {
		for x := 0; x < FIELD_WIDTH; x++ {
			if fieldBuf[y][x] == 1 {
				fmt.Print("[]")
			} else {
				fmt.Print(" _")
			}
		}
		fmt.Println()
	}
}

func (g *Game) FixBlock() {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if BLOCKS[g.block][y][x] == 1 {
				g.field[y+g.pos.y][x+g.pos.x] = 1
			}
		}
	}
}

func (g *Game) EraseLine() {
	for y := 0; y < FIELD_HEIGHT-1; y++ {
		isFilled := true
		for x := 0; x < FIELD_WIDTH-1; x++ {
			if g.field[y][x] == 0 {
				isFilled = false
				break
			}
		}
		if isFilled {
			for y2 := y; y2 > 0; y2-- {
				for x := 0; x < FIELD_WIDTH; x++ {
					g.field[y2][x] = g.field[y2-1][x]
				}
			}
		}
	}
}

func (g *Game) MoveBlock(newPos Position) {
	if !IsCollision(g.field, newPos, g.block) {
		g.pos = newPos
	}
}
