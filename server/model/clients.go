package model

import (
	"github.com/google/uuid"
)

type Client struct {
	Pos Coordinate
	Id  uuid.UUID
}

type ClientList struct {
	Clients  []Client
	Settings *Settings
}

func (cl *ClientList) GetPos(id uuid.UUID) *Coordinate {
	for _, client := range cl.Clients {
		if client.Id == id {
			return &client.Pos
		}
	}
	return nil
}

func (cl *ClientList) GetClientsFromPos(pos Coordinate) []Client {
	list := make([]Client, 0)
	for _, client := range cl.Clients {
		if client.Pos.X == pos.X && client.Pos.Y == pos.Y {
			list = append(list, client)
		}
	}
	return list
}

func (cl *ClientList) GetAvailableSlot(pos Coordinate) *int {
	clientsOnPos := cl.GetClientsFromPos(pos)
	// No space left
	if len(clientsOnPos) >= cl.Settings.Depth {
		return nil
	}

	// Create map with all occupied slots on the z-Coordinate
	occupiedSlots := make(map[int]int)
	for _, c := range clientsOnPos {
		occupiedSlots[*c.Pos.Z] = 1
	}

	// Get the first available position
	for i := 0; i < cl.Settings.Depth; i++ {
		if occupiedSlots[i] != 1 {
			return &i
		}
	}
	return nil
}
