package main

import (
	"github.com/nsf/termbox-go"
	"github.com/taybartski/log"
	"os"
)

func main() {
	log.SetLevel(log.LEVELWARN)

	SetupDisplay()
	defer termbox.Close()

	// bus := make(chan BusMessage) // command bus
	bus := make(chan string) // command bus

	fname := os.Args[1]
	if fname == "" {
		panic("Filename must be provided")
	}
	go DrawLoop(fname, bus)
	HandleInput(bus)
}
