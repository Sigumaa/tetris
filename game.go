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
	field  Field
	pos    Position
	block  BlockShape
	hold   BlockShape
	holded bool
}

func NewGame() *Game {
	g := new(Game)

	g.field = [FIELD_HEIGHT][FIELD_WIDTH]int{
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int(WALL), 0},
		{0, int(WALL), int(WALL), int(WALL), int(WALL), int(WALL), int(WALL), int(WALL), int(WALL), int(WALL), int(WALL), int(WALL), int(WALL), int(WALL), 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	g.pos.Init()
	g.block = BLOCKS[distribution{}.BlockKind()]
	g.hold = NONE_BLOCK
	g.holded = false
	return g
}

func IsCollision(field Field, pos Position, block BlockShape) bool {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if y+pos.y >= FIELD_HEIGHT || x+pos.x >= FIELD_WIDTH {
				continue
			}
			if (block[y][x] != int(NONE)) && (field[y+pos.y][x+pos.x] != int(NONE)) {
				return true
			}
		}
	}
	return false
}

func (g *Game) GhostPos() Position {
	ghostPos := g.pos
	for !IsCollision(g.field, Position{ghostPos.x, ghostPos.y + 1}, g.block) {
		ghostPos.y++
	}
	return ghostPos
}

func (g *Game) Draw() {
	fieldBuf := g.field
	ghostPos := g.GhostPos()

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if g.block[y][x] != int(NONE) {
				fieldBuf[y+ghostPos.y][x+ghostPos.x] = int(GHOST)
			}
		}
	}

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if g.block[y][x] != int(NONE) {
				fieldBuf[y+g.pos.y][x+g.pos.x] = g.block[y][x]
			}
		}
	}

	fmt.Print("\033[2;28HHOLD")
	if g.hold != NONE_BLOCK {
		for y := 0; y < 4; y++ {
			fmt.Printf("\033[%d;%dH", y+3, 28) // カーソルを移動
			for x := 0; x < 4; x++ {
				fmt.Printf("%s", ColorTable[g.hold[y][x]])
			}
			fmt.Println()
		}
	}

	fmt.Print("\x1b[H")
	for y := 0; y < FIELD_HEIGHT-1; y++ {
		for x := 1; x < FIELD_WIDTH-1; x++ {
			fmt.Print(ColorTable[fieldBuf[y][x]])
		}
		fmt.Println()
	}
	fmt.Println("\x1b[0m")
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
			if g.block[y][x] != int(NONE) {
				g.field[y+g.pos.y][x+g.pos.x] = g.block[y][x]
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

func (g *Game) HardDrop() {
	for {
		newPos := Position{g.pos.x, g.pos.y + 1}
		if IsCollision(g.field, newPos, g.block) {
			break
		}
		g.pos = newPos
	}
	g.MoveBlock(g.pos)
}

func (g *Game) Hold() {
	if g.holded {
		return
	}

	if g.hold == NONE_BLOCK {
		g.hold = g.block
		g.SpawnBlock()
	} else {
		g.block, g.hold = g.hold, g.block
		g.pos.Init()
	}
	g.holded = true
}

func (g *Game) landing() error {
	g.FixBlock()
	g.EraseLine()
	if err := g.SpawnBlock(); err != nil {
		return err
	}
	g.holded = false
	return nil
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
	fmt.Print("\x1b[0m\x1b[2J\x1b[?25h")
	os.Exit(0)
}
