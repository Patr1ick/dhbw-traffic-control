package model

type PayloadCoordinates struct {
	Pos Coordinate `form:"pos"`
}

type PayloadMove struct {
	Target Coordinate `form:"target"`
	Id     int        `form:"id"`
}
