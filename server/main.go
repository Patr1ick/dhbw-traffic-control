package main

import (
	"log"
	"os"

	"github.com/Patr1ick/dhbw-traffic-control/server/model"
	"github.com/Patr1ick/dhbw-traffic-control/server/server"
	"github.com/akamensky/argparse"
	"github.com/logrusorgru/aurora/v3"
)

func main() {

	parser := argparse.NewParser("Distributed Systems", "Vehicle Control System")
	w := parser.Int("x", "width", &argparse.Options{Required: true, Help: "Width of the field"})
	h := parser.Int("y", "height", &argparse.Options{Required: true, Help: "Height of the field"})
	d := parser.Int("z", "depth", &argparse.Options{Required: true, Help: "Depth of the field"})

	err := parser.Parse(os.Args)

	if err != nil {
		log.Println(aurora.Red(err.Error()))
		os.Exit(1)
	}

	settings := &model.Settings{
		Width:  *w,
		Height: *h,
		Depth:  *d,
	}

	server.Start(settings)
}
