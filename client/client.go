package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type Pos struct {
	X int
	Y int
	Z int
}

type StartPos struct {
	Pos Pos
	Id  uuid.UUID
}

type TargetPos struct {
	Pos Pos
	Id  uuid.UUID
}

type UpdatePos struct {
	NewPos Pos
	OldPos Pos
}

func main() {

	var xStartPos = 0
	var yStartPos = 0

	var xTargetPos = 13
	var yTargetPos = 12
	//Timestamp start

	//timeBeginn := time.Now()

	var myStartPos StartPos
	myStartPos.Pos.X = xStartPos
	myStartPos.Pos.Y = yStartPos

	var myTargetPos TargetPos
	myTargetPos.Pos.X = xTargetPos
	myTargetPos.Pos.Y = yTargetPos

	var myUpdatePos UpdatePos

	start(&myStartPos)
	//Schleife, bis am ziel angekommen
	myTargetPos.Id = myStartPos.Id

	for {
		move(&myTargetPos, &myUpdatePos)
		fmt.Println("target: ", myTargetPos.Pos.X, " ", myTargetPos.Pos.Y)
		fmt.Println("update: ", myUpdatePos.NewPos.X, " ", myUpdatePos.NewPos.Y)
		if (myUpdatePos.NewPos.X == myTargetPos.Pos.X) &&
			(myUpdatePos.NewPos.Y == myTargetPos.Pos.Y) {
			break
		}
	}

	//Timestamp stop
	//timeEnd := time.Now()
	//timeFinal := timeBeginn.Sub(timeEnd)
}

func start(startPos *StartPos) {
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
	fmt.Println("Response: ", startPos)
}

func move(targetPos *TargetPos, updatePos *UpdatePos) {
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
