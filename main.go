package main

import (
	"github.com/taybartski/log"
	"os"
)

func main() {
	log.SetLevel(log.LEVELWARN)

	SetupDisplay()

	// bus := make(chan BusMessage) // command bus
	bus := make(chan string) // command bus

	fname := os.Args[1]
	if fname == "" {
		panic("Filename must be provided")
	}
	ns := WakeUp()
	go DrawLoop(fname, ns.SpinalCord)
	go HandleInput(ns.SpinalCord)
	ns.AutoNom()
}
