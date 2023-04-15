package model

import "github.com/google/uuid"

type PayloadCoordinates struct {
	Pos Coordinate `form:"pos"`
}

type PayloadMove struct {
	Target Coordinate `form:"target"`
	Id     uuid.UUID  `form:"id"`
}
