package main

import (
	"fmt"
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
			if (field[y+pos.y+1][x+pos.x] & BLOCKS[block][y][x]) == 1 {
				return true
			}
		}
	}
	return false
}

func main() {
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

	fmt.Println("\x1b[2J\x1b[H\x1b[?25l")

	for {
		fieldBuf := field

		if !IsCollision(field, pos, I) {
			pos.y++
		}

		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				if BLOCKS[I][y][x] == 1 {
					fieldBuf[pos.y+y][pos.x+x] = 1
				}
			}
		}

		// フィールドを描画
		fmt.Println("\x1b[H")
		for y := 0; y < FIELD_HEIGHT; y++ {
			for x := 0; x < FIELD_WIDTH; x++ {
				if fieldBuf[y][x] == 1 {
					fmt.Print("[]")
				} else {
					fmt.Print(" .")
				}
			}
			fmt.Println()
		}

		q := false
		for i := 0; i < 60; i++ {
			time.Sleep(1 * time.Millisecond)
			if IsKeyPressed(KEY_Q) {
				q = true
				break
			}
		}
		if q {
			break
		}
	}
	fmt.Println("\x1b[?25h")
}
