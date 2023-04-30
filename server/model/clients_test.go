package model

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetPos(t *testing.T) {
	// Init
	uuid := uuid.New()
	clientList := ClientList{
		Clients: []Client{
			{
				Pos: Coordinate{
					X: 10,
					Y: 10,
				},
				Id: uuid,
			},
		},
		Settings: nil,
	}

	pos := clientList.GetPos(uuid)
	assert.Equal(t, pos.X, 10)
	assert.Equal(t, pos.Y, 10)
}

func TestGetAvailableSlot(t *testing.T) {
	// Init

	clientList := ClientList{
		Clients: []Client{
			{
				Pos: Coordinate{
					X: 10,
					Y: 10,
					Z: new(int),
				},
				Id: uuid.New(),
			},
		},
		Settings: &Settings{
			Depth: 2,
		},
	}

	slot := clientList.GetAvailableSlot(Coordinate{X: 10, Y: 10})
	t.Log(slot)
	assert.Equal(t, *slot, 1)
	assert.NotEqual(t, *slot, nil)
}

func TestGetClientsFromPos(t *testing.T) {
	// Init
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	clientList := ClientList{
		Clients: []Client{
			{
				Pos: Coordinate{
					X: 10,
					Y: 10,
				},
				Id: uuid1,
			},
			{
				Pos: Coordinate{
					X: 10,
					Y: 10,
				},
				Id: uuid2,
			},
		},
		Settings: nil,
	}

	clients := clientList.GetClientsFromPos(Coordinate{X: 10, Y: 10})
	assert.Equal(t, clientList.Clients[0].Id, clients[0].Id)
	assert.Equal(t, clientList.Clients[0].Pos.X, 10)
	assert.Equal(t, clientList.Clients[0].Pos.Y, 10)
	assert.Equal(t, clientList.Clients[1].Id, clients[1].Id)
	assert.Equal(t, clientList.Clients[1].Pos.X, 10)
	assert.Equal(t, clientList.Clients[1].Pos.Y, 10)
	assert.Equal(t, len(clients), 2)
}
