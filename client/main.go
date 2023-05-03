package main

import (
	"math/rand"
	"sync"

	"github.com/Patr1ick/dhbw-traffic-control/client/logic"
)

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			startPos := logic.StartPos{}
			startPos.StartX = rand.Intn(999)
			startPos.StartY = rand.Intn(999)
			startPos.TargetX = rand.Intn(999)
			startPos.TargetY = rand.Intn(999)
			defer wg.Done()
			logic.LeadVehicle(startPos)
		}()
	}
	wg.Wait()

}
