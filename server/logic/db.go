package logic

import (
	"log"

	"github.com/Patr1ick/dhbw-traffic-control/server/model"
	"github.com/gocql/gocql"
)

func LoadTable(session *gocql.Session, settings *model.Settings) (*model.TrafficArea, error) {

	trafficArea := &model.TrafficArea{
		Settings: settings,
	}
	trafficArea.Create()

	scanner := session.Query(`SELECT * FROM traffic_area`).Iter().Scanner()
	for scanner.Next() {
		var (
			x     int
			y     int
			z     int
			value int
		)
		err := scanner.Scan(&x, &y, &z, &value)
		if err != nil {
			return nil, err
		}
		trafficArea.Area[x][y][z] = value
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return trafficArea, nil
}

func SavePos(id int, pos model.Coordinate, session *gocql.Session) error {
	err := session.Query(`INSERT INTO traffic_area (x, y, z, value) VALUES (?, ?, ?, ?)`, pos.X, pos.Y, pos.Z, id).Exec()
	if err != nil {
		return err
	}
	return nil
}
