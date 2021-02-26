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

const (
	Gas = 1 << iota
	EditedGas
	Water
	Sand
)

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
				case Gas:
					pixel = '.'
				case Water:
					pixel = '~'
				case Sand:
					pixel = 's'
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
				if screenSpace[x][y] != 0 {
					switch screenSpace[x][y] {
					/* Sim rules for Sand */
					case Sand:
						if x != 0 && x+1 != width && y+1 != height { /* left limit, right limit, bottom limit */
							if screenSpace[x][y+1] < Sand {
								screenSpace[x][y] = screenSpace[x][y+1]
								screenSpace[x][y+1] = Sand
							} else if screenSpace[x-1][y+1] < Sand {
								screenSpace[x][y] = screenSpace[x-1][y+1]
								screenSpace[x-1][y+1] = Sand
							} else if screenSpace[x+1][y+1] < Sand {
								screenSpace[x][y] = screenSpace[x+1][y+1]
								screenSpace[x+1][y+1] = Sand
							}
						}
					/* Sim rules for Water */
					case Water:
						if x != 0 && x+1 != width && y+1 != height { /* left limit, right limit, bottom limit */
							if screenSpace[x][y+1] < Water {
								screenSpace[x][y] = screenSpace[x][y+1]
								screenSpace[x][y+1] = Water
							} else if screenSpace[x-1][y+1] < Water {
								screenSpace[x][y] = screenSpace[x-1][y+1]
								screenSpace[x-1][y+1] = Water
							} else if screenSpace[x+1][y+1] < Water {
								screenSpace[x][y] = screenSpace[x+1][y+1]
								screenSpace[x+1][y+1] = Water
							} else if screenSpace[x-1][y] < Water {
								screenSpace[x][y] = screenSpace[x-1][y]
								screenSpace[x-1][y] = Water
							} else if screenSpace[x+1][y] < Water {
								screenSpace[x][y] = screenSpace[x+1][y]
								screenSpace[x+1][y] = Water
							}
						}
					/* Sim rules for Gas */
					/* Problem with current implementation -
					 * simulation rule checks happen bottom to top,
					 * So solids and liquids are considered once, but gasses are considered each loop
					 * so they intantly end up at the top
					 * Soln - temporary value EditedGas, to show edited gas, to be skipped in sim rule checks */
					case Gas:
						if x != 0 && x+1 != width && y != 0 { /* left limit, right limit, top limit */
							if screenSpace[x][y-1] < Gas {
								screenSpace[x][y] = screenSpace[x][y-1]
								screenSpace[x][y-1] = EditedGas
							} else if screenSpace[x-1][y-1] < Gas {
								screenSpace[x][y] = screenSpace[x-1][y-1]
								screenSpace[x-1][y-1] = EditedGas
							} else if screenSpace[x+1][y-1] < Gas {
								screenSpace[x][y] = screenSpace[x+1][y-1]
								screenSpace[x+1][y-1] = EditedGas
							} else if screenSpace[x-1][y] < Gas {
								screenSpace[x][y] = screenSpace[x-1][y]
								screenSpace[x-1][y] = EditedGas
							} else if screenSpace[x+1][y] < Gas {
								screenSpace[x][y] = screenSpace[x+1][y]
								screenSpace[x+1][y] = EditedGas
							}
						}
					case EditedGas:
						screenSpace[x][y] = Gas
					default:
					}
				}
			}
		}
		/* Spouting? gas @ 1/2 width from the bottom */
		if screenSpace[width/2][height-2] != Gas {
			screenSpace[width/2][height-1] = Gas
		}
		/* Dripping water @ 2/3 width */
		if screenSpace[2*width/3][1] != Water {
			screenSpace[2*width/3][0] = Water
		}
		/* Pouring sand @ 1/3 width */
		if screenSpace[width/3][1] != Sand {
			screenSpace[width/3][0] = Sand
		}

		spaceChannel <- screenSpace

		time.Sleep(16 * time.Millisecond) /* for approx 60 "frames" on terminal */
	}
}
