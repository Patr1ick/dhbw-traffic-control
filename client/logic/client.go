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

	if start(&startPos) != nil {
		log.Println("Error beim start point")
		return fmt.Errorf("could not set start point")
	}

	//Schleife, bis am ziel angekommen
	targetPos.Id = startPos.Id
	log.Println("ID: ", targetPos.Id)
	for {
		move(&targetPos, &myUpdatePos)

		if (myUpdatePos.NewPos.X == targetPos.Pos.X) &&
			(myUpdatePos.NewPos.Y == targetPos.Pos.Y) {
			break
		}
	}
	//Timestamp stop
	timeEnd := time.Now()
	timeFinal := timeEnd.Sub(timeBeginn)
	log.Println("Time: ", timeFinal, " ID: ", targetPos.Id, " startPosX: ", startPos.Pos.X, " startPosY: ", startPos.Pos.Y,
		" -- ", " targetPosX: ", myUpdatePos.NewPos.X, " targetPosY: ", myUpdatePos.NewPos.Y)

	return nil
}

func start(startPos *VehiclePos) error {
	const serverUrl = "http://localhost:8080/v1/traffic/start"

	requestBodyString := fmt.Sprintf(`
	{
		"pos": {
			"x": %d,
			"y": %d
		} 
	}
	`, startPos.Pos.X, startPos.Pos.Y)

	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}

	requestBody := strings.NewReader(requestBodyString)
	response, err := client.Post(serverUrl, "application/json", requestBody)
	log.Printf("%T %+v", err, err)
	if err != nil || response.StatusCode == 500 {
		defer response.Body.Close()
		for i := 0; i < 5; i++ {
			log.Printf("attemp")
			response, err = http.Post(serverUrl, "application/json", requestBody)

			log.Printf("nach request")
			log.Println(response.StatusCode)
			if err != nil || response.StatusCode == 500 {
				defer response.Body.Close()
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
			//time.Sleep(2 * time.Second)
			log.Printf("ende request")
		}
	}
	log.Printf("kein error")
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
	const moveUrl = "http://localhost:8080/v1/traffic/move"
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
	response, err := http.Post(moveUrl, "application/json", requestBody)

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

	//fmt.Println("target: ", updatePos)
	//fmt.Println("Response: ", string(content))
	return nil
}
