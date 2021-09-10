package main

import "cart/src/w4"

func main() {}

//go:export update
func update() {
	for y := 0; y < 20; y++ {
		for x := 0; x < 20; x++ {
			w4.BlitSub(&tileset[0], x*8, y*8, 8, 8, 0, 0, tilesetWidth, w4.BLIT_2BPP)
		}
	}
}
