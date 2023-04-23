package main

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

func main() {
	const serverUrl = "http://localhost:8080/v1/traffic/start"

	requestBody := strings.NewReader(`
		{
			"pos": {
				"x": 3,
				"y": 4

			} 
		}
	`)
	//Timestamp start
	timeBeginn := time.Now()
	response, err := http.Post(serverUrl, "application/json", requestBody)
	log.Print(response)
	if err != nil {
		log.Print(err)
	}

	//Schleife, bis am ziel angekommen

	//Timestamp stop
	timeEnd := time.Now()
	timeFinal := timeBeginn.Sub(timeEnd)
}

type Pos struct {
	X int
	Y int
	Z int
}

type vehiclePos struct {
	Pos Pos
	Id  uuid.UUID
}

func start(x int, y int) {
	const serverUrl = "http://localhost:8080/v1/traffic/start"

	var cor vehiclePos

	requestBodyString := fmt.Sprintf(`
	{
		"pos": {
			"x": %d,
			"y": %d
		} 
	}
	`, 15, 21)

	requestBody := strings.NewReader(requestBodyString)

	response, _ := http.Post(serverUrl, "application/json", requestBody)
	content, _ := io.ReadAll(response.Body)
	json.Unmarshal(content, &cor)
	fmt.Println("Welcome to hs", cor)

	move(cor.Pos.X, cor.Pos.Y, cor.Id.String())

}

func move(x int, y int, id string) {
	const moveUrl = "http://localhost:8080/v1/traffic/move"
	requestBodyString := fmt.Sprintf(`
	{
		"target": {
			"x": %d,
			"y": %d
		},
		"Id": "%s"
	}
	`, x, y, id)
	fmt.Println(requestBodyString)
	requestBody := strings.NewReader(requestBodyString)
	response, _ := http.Post(moveUrl, "application/json", requestBody)
	content, _ := io.ReadAll(response.Body)
	fmt.Println("Welcome to hs", string(content))
}
