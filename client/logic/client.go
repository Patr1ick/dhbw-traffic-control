package logic

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Pos struct {
	X int
	Y int
	Z int
}

type StartPos struct {
	StartX  int
	StartY  int
	TargetX int
	TargetY int
}

type VehiclePos struct {
	Pos Pos
	Id  uuid.UUID
}

type UpdatePos struct {
	NewPos Pos
	OldPos Pos
}

func LeadVehicle(pos StartPos) error {

	//Timestamp start
	timeBeginn := time.Now()

	startPos := VehiclePos{}
	startPos.Pos.X = pos.StartX
	startPos.Pos.Y = pos.StartY

	targetPos := VehiclePos{}
	targetPos.Pos.X = pos.TargetX
	targetPos.Pos.Y = pos.TargetY

	var myUpdatePos UpdatePos

	if err := start(&startPos); err != nil {
		log.Println(err.Error())
		return fmt.Errorf("could not set start point")
	}

	// Loop until target is reached
	targetPos.Id = startPos.Id
	log.Println("Starting ID: ", targetPos.Id)
	for {
		move(&targetPos, &myUpdatePos)

		if (myUpdatePos.NewPos.X == targetPos.Pos.X) &&
			(myUpdatePos.NewPos.Y == targetPos.Pos.Y) {
			break
		}
	}

	// Stop used time
	timeEnd := time.Now()
	timeFinal := timeEnd.Sub(timeBeginn)
	log.Printf("Used time: %v | ID: %v | Start: (%03d,%03d) | Stop: (%03d,%03d)", timeFinal, targetPos.Id, startPos.Pos.X, startPos.Pos.Y, myUpdatePos.NewPos.X, myUpdatePos.NewPos.Y)

	return nil
}

func start(startPos *VehiclePos) error {
	const url = "http://localhost:8080/v1/traffic/start"

	requestBodyString := fmt.Sprintf(`
	{
		"pos": {
			"x": %v,
			"y": %v
		} 
	}
	`, startPos.Pos.X, startPos.Pos.Y)

	response, err := http.Post(
		url,
		"application/json",
		strings.NewReader(requestBodyString),
	)
	if err != nil || response.StatusCode == 500 {

		for i := 0; i < 5; i++ {
			response, err = http.Post(
				url,
				"application/json",
				strings.NewReader(requestBodyString),
			)

			if err != nil || response.StatusCode == 500 {
				continue
			}

			if response.StatusCode == 200 {
				content, err := io.ReadAll(response.Body)
				if err != nil {
					return err
				}
				if err = json.Unmarshal(content, &startPos); err != nil {
					return err
				}
				return nil
			}
		}
		return fmt.Errorf("could not initialise client")
	}
	content, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(content, &startPos); err != nil {
		return err
	}
	return nil
}

func move(targetPos *VehiclePos, updatePos *UpdatePos) error {
	const url = "http://localhost:8080/v1/traffic/move"
	requestBodyString := fmt.Sprintf(`
	{
		"target": {
			"x": %d,
			"y": %d
		},
		"Id": "%s"
	}
	`, targetPos.Pos.X, targetPos.Pos.Y, targetPos.Id)
	requestBody := strings.NewReader(requestBodyString)
	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		return err
	}

	if response.StatusCode == 200 {
		content, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		if err = json.Unmarshal(content, &updatePos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("could not set start point")
	}

	return nil
}
