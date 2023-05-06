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
			time.Sleep(1 * time.Second)
			mu.Lock()
			newPos := Position{game.pos.x, game.pos.y + 1}
			if !IsCollision(game.field, newPos, game.block) {
				game.pos = newPos
			} else {
				game.FixBlock()
				game.EraseLine()
				if err := game.SpawnBlock(); err != nil {
					game.Over()
					mu.Unlock()
					return
				}
			}
			game.Draw()
			mu.Unlock()
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
		} else if IsKeyPressed(KEY_X) {
			mu.Lock()
			game.RotateRight()
			game.Draw()
			mu.Unlock()
			time.Sleep(50 * time.Millisecond)
		} else if IsKeyPressed(KEY_Z) {
			mu.Lock()
			game.RotateLeft()
			game.Draw()
			mu.Unlock()
			time.Sleep(50 * time.Millisecond)
		} else if IsKeyPressed(KEY_Q) {
			fmt.Println("\x1b[2J\x1b[H\x1b[?25l")
			game.Quit()
		}
		time.Sleep(100 * time.Millisecond)
	}

}
