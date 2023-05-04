package logic

import (
	"testing"

	"github.com/Patr1ick/dhbw-traffic-control/server/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Ptr[T any](v T) *T {
	return &v
}

func TestMove(t *testing.T) {
	t.Run(
		"Move to free field",
		func(t *testing.T) {
			uuid := uuid.New()
			clientList := &model.ClientList{
				Clients: []model.Client{
					{
						Pos: model.Coordinate{
							X: 10,
							Y: 10,
							Z: Ptr(0),
						},
						Id: uuid,
					},
				},
				Settings: &model.Settings{
					Width:  1000,
					Height: 1000,
					Depth:  2,
				},
			}
			_, bestPos, err := Move(clientList, uuid, model.Coordinate{X: 11, Y: 11})
			assert.Nil(t, err)
			assert.Equal(t, bestPos.X, 11)
			assert.Equal(t, bestPos.Y, 11)
		},
	)
	t.Run(
		"Try to move to occupied field",
		func(t *testing.T) {
			uuid := uuid.New()
			clientList := &model.ClientList{
				Clients: []model.Client{
					{
						Pos: model.Coordinate{
							X: 10,
							Y: 10,
							Z: Ptr(0),
						},
						Id: uuid,
					},
					{
						Pos: model.Coordinate{
							X: 11,
							Y: 11,
							Z: Ptr(0),
						},
						Id: uuid,
					},
					{
						Pos: model.Coordinate{
							X: 11,
							Y: 11,
							Z: Ptr(1),
						},
						Id: uuid,
					},
				},
				Settings: &model.Settings{
					Width:  1000,
					Height: 1000,
					Depth:  2,
				},
			}
			_, bestPos, err := Move(clientList, uuid, model.Coordinate{X: 11, Y: 11})
			assert.Nil(t, err)
			assert.Equal(t, bestPos.X, 10)
			assert.Equal(t, bestPos.Y, 11)
		},
	)
}

func TestGetDistance(t *testing.T) {
	t.Run(
		"Same position",
		func(t *testing.T) {
			result := GetDistance(model.Coordinate{X: 0, Y: 0}, model.Coordinate{X: 0, Y: 0})
			assert.Zero(t, result)
		},
	)
	t.Run(
		"Correct distance",
		func(t *testing.T) {
			result := GetDistance(model.Coordinate{X: 0, Y: 0}, model.Coordinate{X: 10, Y: 10})
			assert.Equal(t, result, 14.142135623730951)
		},
	)
}
