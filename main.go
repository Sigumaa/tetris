package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"syscall"
	"time"
)

const (
	KEY_PRESSED = 0x8000
	KEY_Q       = 0x51
	KEY_LEFT    = 0x25
	KEY_RIGHT   = 0x27
	KEY_DOWN    = 0x28
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
)

func IsKeyPressed(keyCode int) bool {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(keyCode))
	return ret&KEY_PRESSED != 0
}

const (
	FIELD_WIDTH  = 13
	FIELD_HEIGHT = 21
)

type Field [FIELD_HEIGHT][FIELD_WIDTH]int
type blockKind int

const (
	I blockKind = iota
	O
	S
	Z
	J
	L
	T
)

func (b blockKind) String() string {
	return [...]string{"I", "O", "S", "Z", "J", "L", "T"}[b]
}

type distribution struct{}

func (d distribution) Intn(n int) int {
	return rand.Intn(n)
}

func (d distribution) BlockKind() blockKind {
	switch d.Intn(7) {
	case 0:
		return I
	case 1:
		return O
	case 2:
		return S
	case 3:
		return Z
	case 4:
		return J
	case 5:
		return L
	default:
		return T
	}
}

type BlockShape [4][4]int

var BLOCKS = map[blockKind]BlockShape{
	I: {
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{0, 0, 0, 0},
	},
	O: {
		{0, 0, 0, 0},
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	},
	S: {
		{0, 0, 0, 0},
		{0, 1, 1, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 0},
	},
	Z: {
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	},
	J: {
		{0, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 1, 1, 0},
		{0, 0, 0, 0},
	},
	L: {
		{0, 0, 0, 0},
		{0, 0, 1, 0},
		{1, 1, 1, 0},
		{0, 0, 0, 0},
	},
	T: {
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 0, 0, 0},
	},
}

type Position struct {
	x int
	y int
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

func Draw(field Field, pos Position, block blockKind) {
	fieldBuf := field
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if BLOCKS[block][y][x] == 1 {
				fieldBuf[y+pos.y][x+pos.x] = 1
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

func main() {
	var mu sync.Mutex

	field := [21][13]int{
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
	pos := Position{4, 0}
	block := distribution{}.BlockKind()

	fmt.Println("\x1b[2J\x1b[H\x1b[?25l")

	mu.Lock()
	Draw(field, pos, block)
	mu.Unlock()

	{
		pos := &pos
		field := &field
		block := &block
		go func() {
			for {
				mu.Lock()
				newPos := Position{pos.x, pos.y + 1}
				if !IsCollision(*field, newPos, *block) {
					*pos = newPos
				} else {
					for y := 0; y < 4; y++ {
						for x := 0; x < 4; x++ {
							if BLOCKS[*block][y][x] == 1 {
								(*field)[y+pos.y][x+pos.x] = 1
							}
						}
					}
					*pos = Position{4, 0}
					*block = distribution{}.BlockKind()
				}
				Draw(*field, *pos, *block)
				mu.Unlock()
				time.Sleep(1 * time.Second)
			}
		}()
	}

	for {
		if IsKeyPressed(KEY_LEFT) {
			mu.Lock()
			newPos := Position{
				int(math.Max(float64(pos.x-1), 0)),
				pos.y,
			}
			if !IsCollision(field, newPos, block) {
				pos = newPos
			}
			Draw(field, pos, block)
			mu.Unlock()
		} else if IsKeyPressed(KEY_RIGHT) {
			mu.Lock()
			newPos := Position{pos.x + 1, pos.y}
			if !IsCollision(field, newPos, block) {
				pos = newPos
			}
			Draw(field, pos, block)
			mu.Unlock()
		} else if IsKeyPressed(KEY_DOWN) {
			mu.Lock()
			newPos := Position{pos.x, pos.y + 1}
			if !IsCollision(field, newPos, block) {
				pos = newPos
			}
			Draw(field, pos, block)
			mu.Unlock()
		} else if IsKeyPressed(KEY_Q) {
			fmt.Println("\x1b[?25h")
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}
