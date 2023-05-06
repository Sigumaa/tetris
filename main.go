package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	game := NewGame()

	fmt.Println("\x1b[2J\x1b[H\x1b[?25l")
	mu.Lock()
	game.Draw()
	mu.Unlock()

	go func() {
		for {
			mu.Lock()
			newPos := Position{game.pos.x, game.pos.y + 1}
			if !IsCollision(game.field, newPos, game.block) {
				game.pos = newPos
			} else {
				game.FixBlock()
				game.EraseLine()
				game.pos.Init()
				game.block = distribution{}.BlockKind()
			}
			game.Draw()
			mu.Unlock()
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		if IsKeyPressed(KEY_LEFT) {
			mu.Lock()
			newPos := Position{
				int(math.Max(float64(game.pos.x-1), 0)),
				game.pos.y,
			}
			game.MoveBlock(newPos)
			game.Draw()
			mu.Unlock()
		} else if IsKeyPressed(KEY_RIGHT) {
			mu.Lock()
			newPos := Position{game.pos.x + 1, game.pos.y}
			game.MoveBlock(newPos)
			game.Draw()
			mu.Unlock()
		} else if IsKeyPressed(KEY_DOWN) {
			mu.Lock()
			newPos := Position{game.pos.x, game.pos.y + 1}
			game.MoveBlock(newPos)
			game.Draw()
			mu.Unlock()
		} else if IsKeyPressed(KEY_Q) {
			fmt.Println("\x1b[?25h")
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}
