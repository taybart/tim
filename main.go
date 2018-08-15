package main

import (
	"github.com/nsf/termbox-go"
	"github.com/taybartski/log"
	"os"
)

// HandleSignals used for load balancing signals
// func HandleSignals(ibus, cbus chan string, quit chan int) {
// for {
// select {
// case sig := <-ibus:
// // log.Info("input: %s", sig)
// if sig == "quit" {
// quit <- 1
// }
// cbus <- sig
// }
// }
// }

// HandleInput used for pulling user input
// func HandleInput(bus chan termbox.Event) {
// fg := termbox.ColorDefault
// bg := termbox.ColorDefault
// loop:
// for {
// termbox.Clear(fg, bg)
// switch ev := termbox.PollEvent(); ev.Type {
// case termbox.EventKey:
// switch ev.Key {
// case termbox.KeyEsc:
// break loop
// default:
// print(2, 2, fg, bg, fmt.Sprintf("%v", ev.N))

// }
// case termbox.EventError:
// panic(ev.Err)
// }
// termbox.Flush()
// }
// }

func main() {
	log.SetLevel(log.LEVELWARN)

	err := termbox.Init()
	check(err)
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	// termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)
	termbox.SetOutputMode(termbox.OutputNormal)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	bus := make(chan string) // input bus
	// ibus := make(chan string) // input bus
	// cbus := make(chan string) // command bus
	// quit := make(chan int) // final bus

	fname := os.Args[1]
	if fname == "" {
		panic("Filename must be provided")
	}
	go HandleBuffer(fname, bus)
	HandleInput(bus)
	// go HandleSignals(ibus, cbus, quit)
}
