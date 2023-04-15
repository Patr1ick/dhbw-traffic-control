package model

type Settings struct {
	Width  int
	Height int
	Depth  int
}

func (settings Settings) Valid(pos Coordinate) bool {

	x := pos.X >= 0 && pos.X < settings.Width
	y := pos.Y >= 0 && pos.Y < settings.Height
	if pos.Z != nil {
		z := *pos.Z >= 0 && *pos.Z < settings.Depth
		return x && y && z
	}
	return x && y
}
