package main

import "syscall"

const (
	KEY_PRESSED = 0x8000
	KEY_Q       = 0x51
	KEY_X       = 0x58
	KEY_Z       = 0x5a
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
