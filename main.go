package main

import (
	"fmt"
)

type Space struct {
	w, h int
	s [][]int
}

func main() {
	width := 30
	height := 30
	screen := make([][]int, width, height)
	screen[0][20] = 1
	screen[1][0] = 2
	// struct Space space = {width, height, screen};
	space := Space{ width, height, screen }
	display(space)
}

func display(space Space) {
	fmt.Printf("\n")
	for _, xline := range space.s {
		for x, c := range xline {
			if x == 0 {
				fmt.Printf("|")
			} else if x == space.w {
				fmt.Printf("|\n")
			}
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
		}
		for x := 0; x < space.w; x++ {
			fmt.Printf("-")
		}
	}
	fmt.Printf("\n")
}

// void sim(struct Space *s)
// {
// 	int width = s->w, height = s->h;
// 	while(1) {
// 		for (int i = 0; i < width*height; i++) {
// 			if (1 == s->space[i]) {
// 				if (s->space[i+width] != 1) {
// 					s->space[i] = 0;
// 					s->space[i+width] = 1;
// 				} else if (s->space[i+width-1] != 1) {
// 					s->space[i] = 0;
// 					s->space[i+width-1] = 1;
// 				} else if (s->space[i+width+1] != 1) {
// 					s->space[i] = 0;
// 					s->space[i+width+1] = 1;
// 				}
// 			}
// 		}

// 		if (s->space[width/2 + width] != 1) s->space[width/2] = 1;

// 		display(s);
// 		sleep(1);
// 	}
// 	return;
// }
