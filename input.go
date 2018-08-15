package main

import (
	"github.com/nsf/termbox-go"
)

// HandleInput used for pulling user input
func HandleInput(bus chan string) {
	data := make([]byte, 0, 64)
inputloop:
	for {
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
			if string(data) == "q" {
				break inputloop
			}
			bus <- string(data)

			for {
				ev := termbox.ParseEvent(data)
				if ev.N == 0 {
					break
				}
				copy(data, data[ev.N:])
				data = data[:len(data)-ev.N]
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
