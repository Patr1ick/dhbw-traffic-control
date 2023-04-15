package main

import (
	"log"
	"net/http"
	"strings"
	"time"
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
