package main

import "math/rand"

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
