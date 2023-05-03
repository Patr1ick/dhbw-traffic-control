package main

import (
	"log"
	"math/rand"
	"sync"

	"github.com/Patr1ick/dhbw-traffic-control/client/logic"
)

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 1; i++ {
		wg.Add(1)

		go func() {
			log.Println("Hello")
			/*
				startPos := logic.StartPos{}
				startPos.StartX = rand.Intn(999)
				startPos.StartY = rand.Intn(999)
				startPos.TargetX = rand.Intn(999)
				startPos.TargetY = rand.Intn(999)
			*/
			startPos := logic.StartPos{}
			startPos.StartX = 10
			startPos.StartY = 10
			startPos.TargetX = rand.Intn(11)
			startPos.TargetY = rand.Intn(11)
			defer wg.Done()
			logic.LeadVehicle(startPos)
		}()
	}
	wg.Wait()

}
