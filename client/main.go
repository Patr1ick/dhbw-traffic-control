package main

import (
	"log"
	"math/rand"
	"os"
	"sync"

	"github.com/Patr1ick/dhbw-traffic-control/client/logic"
	"github.com/akamensky/argparse"
)

func main() {

	parser := argparse.NewParser("Distributed Systems", "Vehicle Control System")
	vehicles := parser.Int("x", "vehicles", &argparse.Options{Required: true, Help: "Number of vehicles"})
	address := parser.String("c", "caddy", &argparse.Options{Required: true, Help: "The address to caddy. Should be: localhost:<port> or <ip>:<port>"})
	err := parser.Parse(os.Args)

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	if *vehicles > 1 && *vehicles < 1000 {
		var wg sync.WaitGroup
		for i := 0; i < *vehicles; i++ {
			wg.Add(1)

			go func() {
				startPos := logic.StartPos{}
				startPos.StartX = rand.Intn(999)
				startPos.StartY = rand.Intn(999)
				startPos.TargetX = rand.Intn(999)
				startPos.TargetY = rand.Intn(999)
				defer wg.Done()
				logic.LeadVehicle(startPos, *address)
			}()
		}
		wg.Wait()
	} else {
		log.Println("The number of vehicles must not be less than 1 and not more than 1000")
	}

}
