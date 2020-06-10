package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"goinvaders/invaders"
	"goinvaders/machine"
)

func main() {
	cabinet := machine.NewCabinet()
	game := invaders.NewInvaders(cabinet)

	timer := sdl.GetTicks()
	for !game.ShouldQuit() {
		game.PollEvents()
		if float64(sdl.GetTicks()-timer) > (1.0/machine.FPS)*1000 {
			timer = sdl.GetTicks()
			game.Update()
			game.GpuUpdate()
		}
	}
}
