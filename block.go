package main

import "math/rand"

type BlockKind int

const (
	I BlockKind = iota
	O
	S
	Z
	J
	L
	T
)

type distribution struct{}

func (d distribution) Intn(n int) int {
	return rand.Intn(n)
}

func (d distribution) BlockKind() BlockKind {
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

type BlockColor int

const (
	NONE BlockColor = iota
	WALL
	I_COLOR
	O_COLOR
	S_COLOR
	Z_COLOR
	J_COLOR
	L_COLOR
	T_COLOR
)

var ColorTable = []string{
	"\x1b[48;2;000;000;000m  ",
	"\x1b[48;2;127;127;127m__",
	"\x1b[48;2;000;000;255m__",
	"\x1b[48;2;000;255;000m__",
	"\x1b[48;2;000;255;255m__",
	"\x1b[48;2;255;000;000m__",
	"\x1b[48;2;255;000;255m__",
	"\x1b[48;2;255;127;000m__",
	"\x1b[48;2;255;255;000m__",
}

type BlockShape [4][4]int

var BLOCKS = map[BlockKind]BlockShape{
	I: {
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{int(I_COLOR), int(I_COLOR), int(I_COLOR), int(I_COLOR)},
		{0, 0, 0, 0},
	},
	O: {
		{0, 0, 0, 0},
		{0, int(O_COLOR), int(O_COLOR), 0},
		{0, int(O_COLOR), int(O_COLOR), 0},
		{0, 0, 0, 0},
	},
	S: {
		{0, 0, 0, 0},
		{0, int(S_COLOR), int(S_COLOR), 0},
		{int(S_COLOR), int(S_COLOR), 0, 0},
		{0, 0, 0, 0},
	},
	Z: {
		{0, 0, 0, 0},
		{int(Z_COLOR), int(Z_COLOR), 0, 0},
		{0, int(Z_COLOR), int(Z_COLOR), 0},
		{0, 0, 0, 0},
	},
	J: {
		{0, 0, 0, 0},
		{int(J_COLOR), 0, 0, 0},
		{int(J_COLOR), int(J_COLOR), int(J_COLOR), 0},
		{0, 0, 0, 0},
	},
	L: {
		{0, 0, 0, 0},
		{0, 0, int(L_COLOR), 0},
		{int(L_COLOR), int(L_COLOR), int(L_COLOR), 0},
		{0, 0, 0, 0},
	},
	T: {
		{0, 0, 0, 0},
		{0, int(T_COLOR), 0, 0},
		{int(T_COLOR), int(T_COLOR), int(T_COLOR), 0},
		{0, 0, 0, 0},
	},
}
