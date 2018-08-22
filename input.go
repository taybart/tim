package main

import (
	"github.com/nsf/termbox-go"
	// "github.com/taybartski/log"
)

// HandleInput used for pulling user input
func HandleInput(bus chan interface{}) {
	data := make([]byte, 0, 64)
inputloop:
	for {
		select {
		case ns := <-bus:
			if ns.(string) == "q" {
				break inputloop
			}
		default:
			if cap(data)-len(data) < 32 {
				newdata := make([]byte, len(data), len(data)+32)
				copy(newdata, data)
				data = newdata
			}
			beg := len(data)
			d := data[beg : beg+32]
			switch ev := termbox.PollRawEvent(d); ev.Type {
			case termbox.EventRaw:
				data = data[:beg+ev.N]
				bus <- string(data)
				// if string(data) == "q" {
				// termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				// break inputloop
				// }

				for {
					ev := termbox.ParseEvent(data)
					if ev.N == 0 {
						break
					}
					copy(data, data[ev.N:])
					data = data[:len(data)-ev.N]
				}
			case termbox.EventResize:
				bus <- "resize"
			case termbox.EventError:
				panic(ev.Err)
			}
		}
	}
}
