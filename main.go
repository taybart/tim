package main

import (
	"github.com/nsf/termbox-go"
	"github.com/taybartski/log"
	"os"
)

func main() {
	log.SetLevel(log.LEVELWARN)

	err := termbox.Init()
	check(err)
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.SetOutputMode(termbox.OutputNormal)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	bus := make(chan string) // command bus

	fname := os.Args[1]
	if fname == "" {
		panic("Filename must be provided")
	}
	go HandleBuffer(fname, bus)
	HandleInput(bus)
}
