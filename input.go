package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"strings"
)

var current string
var curev termbox.Event

/* func mouseButtonToStr(k termbox.Key) string {
	switch k {
	case termbox.MouseLeft:
		return "MouseLeft"
	case termbox.MouseMiddle:
		return "MouseMiddle"
	case termbox.MouseRight:
		return "MouseRight"
	case termbox.MouseRelease:
		return "MouseRelease"
	case termbox.MouseWheelUp:
		return "MouseWheelUp"
	case termbox.MouseWheelDown:
		return "MouseWheelDown"
	}
	return "Key"
} */

func modStr(m termbox.Modifier) string {
	var out []string
	if m&termbox.ModAlt != 0 {
		out = append(out, "ModAlt")
	}
	if m&termbox.ModMotion != 0 {
		out = append(out, "ModMotion")
	}
	return strings.Join(out, " | ")
}

func draw() {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)
	print(0, 0, termbox.ColorMagenta, coldef, "Press 'q' to quit")
	print(0, 1, coldef, coldef, current)
	switch curev.Type {
	case termbox.EventKey:
		print(0, 2, coldef, coldef,
			fmt.Sprintf("EventKey: k: %d, c: %c, mod: %s", curev.Key, curev.Ch, modStr(curev.Mod)))
	/* case termbox.EventMouse:
	print(0, 2, coldef, coldef,
		fmt.Sprintf("EventMouse: x: %d, y: %d, b: %s, mod: %s",
			curev.MouseX, curev.MouseY, mouseButtonToStr(curev.Key), modStr(curev.Mod))) */
	case termbox.EventNone:
		print(0, 2, coldef, coldef, "EventNone")
	}
	print(0, 3, coldef, coldef, fmt.Sprintf("%d", curev.N))
	termbox.Flush()
}

// HandleInput used for pulling user input
func HandleInput(bus chan string) {
	// draw()
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
			current = fmt.Sprintf("%q", data)
			if current == `"q"` {
				break inputloop
			}
			bus <- string(data)

			for {
				ev := termbox.ParseEvent(data)
				if ev.N == 0 {
					break
				}
				curev = ev
				copy(data, data[curev.N:])
				data = data[:len(data)-curev.N]
			}
		case termbox.EventError:
			panic(ev.Err)
		}
		// draw()
	}
}
