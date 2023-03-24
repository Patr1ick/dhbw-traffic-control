package model

type Coordinate struct {
	X int `form:"x"`
	Y int `form:"y"`
	Z *int
}
