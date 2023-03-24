package model

import (
	"fmt"

	"github.com/logrusorgru/aurora/v3"
)

type TrafficArea struct {
	Area     [][][]int
	Width    int
	Height   int
	Depth    int
	Settings *Settings
}

func (ta *TrafficArea) Create() {

	var area [][][]int = make([][][]int, ta.Settings.Width)
	for i := range area {
		area[i] = make([][]int, ta.Settings.Height)

		for j := range area[i] {
			area[i][j] = make([]int, ta.Settings.Depth)

			for k := 0; k < ta.Settings.Depth; k++ {
				area[i][j][k] = -1
			}
		}
	}

	ta.Area = area
	ta.Width = ta.Settings.Width
	ta.Height = ta.Settings.Height
	ta.Depth = ta.Settings.Depth
}

func (ta *TrafficArea) Set(id int, pos Coordinate) (*Coordinate, error) {
	freePos := -1
	for i := 0; i < len(ta.Area[pos.X][pos.Y]); i++ {
		if ta.Area[pos.X][pos.Y][i] == id {
			return nil, fmt.Errorf("already in position")
		}
		if ta.Area[pos.X][pos.Y][i] == -1 {
			freePos = i
		}
	}

	if freePos == -1 {
		return nil, fmt.Errorf("field not empty")
	}

	ta.Area[pos.X][pos.Y][freePos] = id
	return &Coordinate{X: pos.X, Y: pos.Y, Z: &freePos}, nil
}

func (ta *TrafficArea) Remove(id int, pos Coordinate) error {
	for i := 0; i < len(ta.Area[pos.X][pos.Y]); i++ {
		if ta.Area[pos.X][pos.Y][i] == id {
			ta.Area[pos.X][pos.Y][i] = -1
			return nil
		}
	}
	return fmt.Errorf("not found")

}

func (ta *TrafficArea) Get(id int) *Coordinate {
	for x := 0; x < ta.Width; x++ {
		for y := 0; y < ta.Height; y++ {
			for d := 0; d < ta.Depth; d++ {
				if ta.Area[x][y][d] == id {
					return &Coordinate{
						X: x,
						Y: y,
						Z: &d,
					}
				}
			}
		}
	}
	return nil
}

func (ta *TrafficArea) IsFree(pos Coordinate) int {
	if pos.X > ta.Settings.Width || pos.X < 0 {
		return -1
	}
	if pos.Y > ta.Settings.Height || pos.Y < 0 {
		return -1
	}

	for i := 0; i < ta.Settings.Depth; i++ {
		if ta.Area[pos.X][pos.Y][i] == -1 {
			return i
		}
	}
	return -1
}

func (ta *TrafficArea) Print() {
	fmt.Printf("\nTraffic Area (X: %d, Y: %d, d: %d):\n", ta.Width, ta.Height, ta.Depth)
	for i := range ta.Area {
		fmt.Print("|")
		for j := range ta.Area[i] {
			for _, z := range ta.Area[i][j] {
				if z == -1 {
					fmt.Printf("%3d", aurora.Red(z))
				} else {
					fmt.Printf("%3d", aurora.Green(z))
				}
			}
			fmt.Print(" |")
		}
		fmt.Print("\n")
	}
}

func (ta *TrafficArea) ToTable() ([]int, []int, []int, []int) {
	numElem := ta.Width * ta.Height * ta.Depth

	X := make([]int, numElem)
	Y := make([]int, numElem)
	Z := make([]int, numElem)
	value := make([]int, numElem)

	i := 0

	for x := range ta.Area {
		for y := range ta.Area[x] {
			for z := range ta.Area[x][y] {
				X[i] = x
				Y[i] = y
				Z[i] = z
				value[i] = ta.Area[x][y][z]
				i++
			}
		}
	}
	return X, Y, Z, value
}

func (ta *TrafficArea) ToArray(X []int, Y []int, Z []int, value []int) {
	ta.Create()
	for i, val := range value {
		ta.Area[X[i]][Y[i]][Z[i]] = val
	}
}
