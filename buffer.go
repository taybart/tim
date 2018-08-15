package main

import (
	"github.com/nsf/termbox-go"
	"io/ioutil"
)

// HandleBuffer Used for navigating buffers
func HandleBuffer(fileName string, bus chan string) {
	file, err := ioutil.ReadFile(fileName)
	check(err)

	width, height := termbox.Size()
	d := Display{1, 1, width, height, file}
	for {
		d.Draw()
		termbox.Flush()
		select {
		case s := <-bus:
			switch s {
			case "h":
				d.col--
				if d.col < 0 {
					d.col = 0
				}
			case "j":
				d.row++
				if d.row > d.height {
					d.row = d.height - 1
				}
			case "k":
				d.row--
				if d.row < 0 {
					d.row = 0
				}
			case "l":
				d.col++
				if d.col > d.width {
					d.col = d.width - 1
				}
			}
		}
	}
}
