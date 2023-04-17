package logic

import (
	"fmt"
	"math"

	"github.com/Patr1ick/dhbw-traffic-control/server/model"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

func InitClient(session *gocql.Session, cl *model.ClientList, newClient *model.Client) error {

	z := cl.GetAvailableSlot(newClient.Pos)

	if z == nil {
		return fmt.Errorf("position not free")
	}

	newClient.Pos.Z = z

	if err := AddClient(*newClient, session); err != nil {
		return err
	}
	return nil
}

func Move(cl *model.ClientList, id uuid.UUID, target model.Coordinate) (*model.Coordinate, *model.Coordinate, error) {
	currentPos := cl.GetPos(id)
	if currentPos == nil {
		return nil, nil, fmt.Errorf("not initialised")
	}
	bestPos := currentPos
	distance := GetDistance(*currentPos, target)

	for xOffset := -1; xOffset <= 1; xOffset++ {
		for yOffset := -1; yOffset <= 1; yOffset++ {

			newPos := model.Coordinate{
				X: currentPos.X + xOffset,
				Y: currentPos.Y + yOffset,
			}
			if newPos.X < 0 {
				newPos.X = 0
			}
			if newPos.X > cl.Settings.Width-1 {
				newPos.X = cl.Settings.Width - 1
			}
			if newPos.Y < 0 {
				newPos.Y = 0
			}
			if newPos.Y > cl.Settings.Height-1 {
				newPos.Y = cl.Settings.Height - 1
			}
			if z := cl.GetAvailableSlot(newPos); z != nil {
				newPos.Z = z
				newDistance := GetDistance(newPos, target)
				if newDistance < distance {
					distance = newDistance
					bestPos = &newPos
				}
			}
		}
	}
	return currentPos, bestPos, nil
}

func GetDistance(pos1 model.Coordinate, pos2 model.Coordinate) float64 {
	if pos1.X == pos2.X && pos1.Y == pos2.Y {
		return 0
	}
	return math.Sqrt(math.Pow(float64(pos1.X)-float64(pos2.X), 2) + math.Pow(float64(pos1.Y)-float64(pos2.Y), 2))
}
