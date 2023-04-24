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

func LeadVehicle(pos StartPos) {

	//Timestamp start
	timeBeginn := time.Now()

	startPos := VehiclePos{}
	startPos.Pos.X = pos.StartX
	startPos.Pos.Y = pos.StartY

	targetPos := VehiclePos{}
	targetPos.Pos.X = pos.TargetX
	targetPos.Pos.Y = pos.TargetY

	var myUpdatePos UpdatePos

	start(&startPos)
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
}

func start(startPos *VehiclePos) {
	const serverUrl = "http://localhost:8080/v1/traffic/start"

	requestBodyString := fmt.Sprintf(`
	{
		"pos": {
			"x": %d,
			"y": %d
		} 
	}
	`, startPos.Pos.X, startPos.Pos.Y)

	requestBody := strings.NewReader(requestBodyString)
	response, _ := http.Post(serverUrl, "application/json", requestBody)
	content, _ := io.ReadAll(response.Body)
	json.Unmarshal(content, &startPos)
}

func move(targetPos *VehiclePos, updatePos *UpdatePos) {
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
	response, _ := http.Post(moveUrl, "application/json", requestBody)
	content, _ := io.ReadAll(response.Body)
	json.Unmarshal(content, &updatePos)
	//fmt.Println("target: ", updatePos)
	//fmt.Println("Response: ", string(content))
}
