package main

import (
	"fmt"
	"time"
)

type Space struct {
	w, h int
	s    [][]int
}

func main() {
	width := 30
	height := 30
	screen := make([][]int, width, height)
	for i := range screen {
		for y := 0; y < height; y++ {
			screen[i] = append(screen[i], 0)
		}
	}
	space := Space{
		w: width,
		h: height,
		s: screen}
	sim(space)
}

func display(space Space) {
	fmt.Printf("\n")
	for x := 0; x < space.w; x++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
	for _, xline := range space.s {
		for x, c := range xline {
			var pixel rune
			switch c {
			case 1:
				pixel = 's'
			case 2:
				pixel = 'w'
			default:
				pixel = ' '
			}
			fmt.Printf("%c", pixel)
			if x == space.w-1 {
				fmt.Printf("|\n")
			}
		}
	}
	for x := 0; x < space.w; x++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
}

func sim(space Space) {
	width := space.w
	height := space.h
	for i := 0; i < 100; i++ {
		for x := width - 1; x >= 0; x-- {
			for y := height - 1; y >= 0; y-- {
				if x+1 != width && y+1 != height {
					/* Sim for Sand */
					if 1 == space.s[x][y] {
						if space.s[x][y+1] != 1 {
							space.s[x][y] = space.s[x][y+1]
							space.s[x][y+1] = 1
						} else if space.s[x-1][y+1] != 1 {
							space.s[x][y] = space.s[x-1][y+1]
							space.s[x-1][y+1] = 1
						} else if space.s[x+1][y+1] != 1 {
							space.s[x][y] = space.s[x+1][y+1]
							space.s[x+1][y+1] = 1
						}
					}
					/* Sim for Sand */
					if 2 == space.s[x][y] {
						if space.s[x][y+1] == 0 {
							space.s[x][y] = space.s[x][y+1]
							space.s[x][y+1] = 2
						} else if space.s[x-1][y+1] == 0 {
							space.s[x][y] = space.s[x-1][y+1]
							space.s[x-1][y+1] = 2
						} else if space.s[x+1][y+1] == 0 {
							space.s[x][y] = space.s[x+1][y+1]
							space.s[x+1][y+1] = 2
						} else if space.s[x-1][y] == 0 {
							space.s[x][y] = space.s[x-1][y]
							space.s[x-1][y] = 2
						} else if space.s[x+1][y] == 0 {
							space.s[x][y] = space.s[x+1][y]
							space.s[x+1][y] = 2
						}
					}

				}
			}
		}
		if space.s[width/3][1] != 1 {
			space.s[width/3][0] = 1
		}
		if space.s[2*width/3][1] != 2 {
			space.s[2*width/3][0] = 2
		}
		display(space)

		time.Sleep(50 * time.Millisecond)
	}
}
