package logic

import (
	"fmt"
	"log"

	"github.com/Patr1ick/dhbw-traffic-control/server/model"
	"github.com/google/uuid"
	"github.com/yugabyte/gocql"
)

func LoadTable(session *gocql.Session, settings *model.Settings) (*model.ClientList, error) {

	cl := &model.ClientList{
		Settings: settings,
	}

	clients := make([]model.Client, 0)

	scanner := session.Query(`SELECT x, y, z, id FROM traffic_control.clients`).Iter().Scanner()
	for scanner.Next() {
		var (
			x  int
			y  int
			z  int
			id gocql.UUID
		)
		err := scanner.Scan(&x, &y, &z, &id)
		if err != nil {
			return nil, err
		}

		uuid, err := uuid.Parse(id.String())
		if err != nil {
			return nil, fmt.Errorf("could not convert uuid")
		}

		clients = append(clients, model.Client{
			Pos: model.Coordinate{
				X: x,
				Y: y,
				Z: &z,
			},
			Id: uuid,
		})

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	cl.Clients = clients

	return cl, nil
}

func AddClient(client model.Client, session *gocql.Session) error {
	uuid, err := gocql.ParseUUID(client.Id.String())
	if err != nil {
		return err
	}
	err = session.Query(`INSERT INTO traffic_control.clients (x, y, z, id) VALUES (?, ?, ?, ?)`, client.Pos.X, client.Pos.Y, client.Pos.Z, uuid).Exec()
	if err != nil {
		return err
	}
	return nil
}
func UpdateClient(client model.Client, session *gocql.Session) error {
	uuid, err := gocql.ParseUUID(client.Id.String())
	if err != nil {
		return err
	}
	err = session.Query(`UPDATE traffic_control.clients SET x = ?, y = ?, z = ? WHERE id = ?`, client.Pos.X, client.Pos.Y, client.Pos.Z, uuid).Exec()
	if err != nil {
		return err
	}
	return nil
}
