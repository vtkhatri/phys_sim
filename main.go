package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
	//"github.com/gdamore/tcell/v2"
)

func main() {
	width := 50
	height := 35

	spaceChannel := make(chan [][]int)
	go display(spaceChannel)
	sim(width, height, spaceChannel)
}

func display(spaceChannel chan [][]int) {
	for space := range spaceChannel {
		switch runtime.GOOS {
		case "windows":
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
		case "linux":
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		default:
		}

		/* Starting printing */
		for y := range space[0] {
			for x := range space {
				var pixel rune
				switch space[x][y] {
				case 1:
					pixel = 's'
				case 2:
					pixel = '~'
				default:
					pixel = ' '
				}
				fmt.Printf("%c", pixel)
			}
			fmt.Printf("\n")
		}
	}
}

func sim(width int, height int, spaceChannel chan [][]int) {
	screenSpace := make([][]int, width)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			screenSpace[x] = append(screenSpace[x], 0)
		}
	}
	for {
		for x := width - 1; x >= 0; x-- {
			for y := height - 1; y >= 0; y-- {
				if screenSpace[x][y] != 0 && x != 0 && x+1 != width && y+1 != height {
					switch screenSpace[x][y] {
					/* Sim rules for Sand */
					case 1:
						if screenSpace[x][y+1] != 1 {
							screenSpace[x][y] = screenSpace[x][y+1]
							screenSpace[x][y+1] = 1
						} else if screenSpace[x-1][y+1] != 1 {
							screenSpace[x][y] = screenSpace[x-1][y+1]
							screenSpace[x-1][y+1] = 1
						} else if screenSpace[x+1][y+1] != 1 {
							screenSpace[x][y] = screenSpace[x+1][y+1]
							screenSpace[x+1][y+1] = 1
						}
					/* Sim rules for Water */
					case 2:
						if screenSpace[x][y+1] == 0 {
							screenSpace[x][y] = screenSpace[x][y+1]
							screenSpace[x][y+1] = 2
						} else if screenSpace[x-1][y+1] == 0 {
							screenSpace[x][y] = screenSpace[x-1][y+1]
							screenSpace[x-1][y+1] = 2
						} else if screenSpace[x+1][y+1] == 0 {
							screenSpace[x][y] = screenSpace[x+1][y+1]
							screenSpace[x+1][y+1] = 2
						} else if screenSpace[x-1][y] == 0 {
							screenSpace[x][y] = screenSpace[x-1][y]
							screenSpace[x-1][y] = 2
						} else if screenSpace[x+1][y] == 0 {
							screenSpace[x][y] = screenSpace[x+1][y]
							screenSpace[x+1][y] = 2
						}
					default:
					}
				}
			}
		}
		/* Pouring sand @ 1/3 width */
		if screenSpace[width/3][1] != 1 {
			screenSpace[width/3][0] = 1
		}
		/* Dripping water @ 2/3 width */
		if screenSpace[2*width/3][1] != 2 {
			screenSpace[2*width/3][0] = 2
		}
		spaceChannel <- screenSpace

		time.Sleep(16 * time.Millisecond) /* for approx 60 "frames" on terminal */
	}
}
