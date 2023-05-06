package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	FIELD_WIDTH  = 15
	FIELD_HEIGHT = 22
)

type Field [FIELD_HEIGHT][FIELD_WIDTH]int

type Position struct {
	x int
	y int
}

func (p *Position) Init() {
	p.x = 5
	p.y = 0
}

type Game struct {
	field Field
	pos   Position
	block BlockShape
}

func NewGame() *Game {
	g := new(Game)

	g.field = [FIELD_HEIGHT][FIELD_WIDTH]int{
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	g.pos.Init()
	g.block = BLOCKS[distribution{}.BlockKind()]
	return g
}

func IsCollision(field Field, pos Position, block BlockShape) bool {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if y+pos.y >= FIELD_HEIGHT || x+pos.x >= FIELD_WIDTH {
				continue
			}
			if (field[y+pos.y][x+pos.x] & block[y][x]) == 1 {
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
			if g.block[y][x] == 1 {
				fieldBuf[y+g.pos.y][x+g.pos.x] = 1
			}
		}
	}

	fmt.Print("\x1b[H")
	for y := 0; y < FIELD_HEIGHT-1; y++ {
		for x := 1; x < FIELD_WIDTH-1; x++ {
			if fieldBuf[y][x] == 1 {
				fmt.Print("[]")
			} else {
				fmt.Print(" _")
			}
		}
		fmt.Println()
	}
}

func (g *Game) SpawnBlock() error {
	g.pos.Init()
	g.block = BLOCKS[distribution{}.BlockKind()]
	if IsCollision(g.field, g.pos, g.block) {
		return errors.New("game over")
	}
	return nil
}

func (g *Game) FixBlock() {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if g.block[y][x] == 1 {
				g.field[y+g.pos.y][x+g.pos.x] = 1
			}
		}
	}
}

func (g *Game) EraseLine() {
	for y := 1; y < FIELD_HEIGHT-2; y++ {
		isFilled := true
		for x := 2; x < FIELD_WIDTH-2; x++ {
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

func (g *Game) RotateRight() {
	rotated := BlockShape{}
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			rotated[y][x] = g.block[3-x][y]
		}
	}
	if !IsCollision(g.field, g.pos, rotated) {
		g.block = rotated
	}
}

func (g *Game) RotateLeft() {
	rotated := BlockShape{}
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			rotated[y][x] = g.block[x][3-y]
		}
	}
	if !IsCollision(g.field, g.pos, rotated) {
		g.block = rotated
	}
}

func (g *Game) Over() {
	g.Draw()
	fmt.Println("Game Over")
	fmt.Println("press `q` key to exit")
}

func (g *Game) Quit() {
	fmt.Print("\x1b[?25h")
	os.Exit(0)
}
