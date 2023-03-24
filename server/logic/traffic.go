package logic

import (
	"fmt"
	"math"

	"github.com/Patr1ick/dhbw-traffic-control/server/model"
)

func Start(trafficArea *model.TrafficArea, id int) (*model.Coordinate, error) {
	currentPos := trafficArea.Get(id)
	if currentPos == nil {
		return currentPos, fmt.Errorf("client already exists")
	}

	for y := 0; y < len(trafficArea.Area[0]); y++ {
		pos := model.Coordinate{X: 0, Y: y}
		if trafficArea.IsFree(pos) != -1 {
			trafficArea.Set(id, pos)
			return &pos, nil
		}
	}

	return nil, fmt.Errorf("no free position found")
}

func Move(trafficArea *model.TrafficArea, id int, target model.Coordinate) (*model.Coordinate, *model.Coordinate, error) {
	currentPos := trafficArea.Get(id)
	bestPos := currentPos
	distance := GetDistance(*currentPos, target)
	for xOffset := -1; xOffset <= 1; xOffset++ {
		for yOffset := -1; yOffset <= 1; yOffset++ {
			x := currentPos.X + xOffset
			y := currentPos.Y + yOffset

			//  Check boundaries
			if x < 0 {
				x = 0
			}
			if x > trafficArea.Width {
				x = trafficArea.Width
			}
			if y < 0 {
				x = 0
			}
			if y > trafficArea.Height {
				y = trafficArea.Height
			}

			newPos := model.Coordinate{X: x, Y: y}
			if z := trafficArea.IsFree(newPos); z != -1 {
				newPos.Z = &z
				newDistance := GetDistance(newPos, target)
				if newDistance < distance {
					distance = newDistance
					bestPos = &newPos
				}
			}
		}
	}

	// Move
	if err := trafficArea.Remove(id, *currentPos); err != nil {
		return nil, nil, err
	}
	if _, err := trafficArea.Set(id, *bestPos); err != nil {
		return nil, nil, err
	}
	return currentPos, bestPos, nil
}

func GetDistance(pos1 model.Coordinate, pos2 model.Coordinate) float64 {
	if pos1.X == pos2.X && pos1.Y == pos2.Y {
		return 0
	}
	return math.Sqrt(math.Pow(float64(pos1.X)-float64(pos2.X), 2) + math.Pow(float64(pos1.Y)-float64(pos2.Y), 2))
}
