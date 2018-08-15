package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"io/ioutil"
)

type Display struct {
	row    int
	col    int
	width  int
	height int
	m      margin
}
type margin struct {
	r int
	t int
	l int
	b int
}

func (d Display) Draw(b Buffer) {
	d.DrawBuffer(b)
	d.DrawMode(b)
	d.DrawCursor()
}

// DrawMode displays current modality
func (d Display) DrawMode(b Buffer) {
	// lines := len(b.data)/d.width + 1
	bar := fmt.Sprintf("buf[%d]:%s (%d,%d), len(%d), width(%d)",
		b.index, string(b.data[b.index]), d.row, d.col, len(b.data), d.width)
	draw(d.width-len(bar), d.height-1, termbox.ColorRed, termbox.ColorDefault, bar)
}

// DrawBuffer displays current file
func (d Display) DrawBuffer(b Buffer) {
	editWidth := d.width - d.m.l - d.m.r
	lines := (len(b.data) / editWidth) + 1

	for i := 0; i < lines; i++ {
		draw(0, i+d.m.t, termbox.ColorYellow, termbox.ColorDefault, fmt.Sprintf("%d", i))
		draw(d.m.l, i+d.m.t, termbox.ColorDefault, termbox.ColorDefault, string(b.data[i*editWidth:(i+1)*editWidth]))
	}
}
func (d Display) DrawCursor() {
	termbox.SetCursor(d.col, d.row)
}

func draw(x, y int, fg, bg termbox.Attribute, s string) {
	for _, c := range s {
		termbox.SetCell(x, y, rune(c), fg, bg)
		x++
	}
}

// DrawLoop used for drawing
func DrawLoop(fileName string, bus chan string) {
	file, err := ioutil.ReadFile(fileName)
	check(err)
	m := margin{1, 1, 3, 1}
	b := Buffer{file, 0, "normal"}

	width, height := termbox.Size()
	d := Display{m.t, m.l, width, height, m}
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		d.Draw(b)
		termbox.Flush()
		select {
		case s := <-bus:
			switch s {
			case "h":
				d.col = decrement(d.col, d.m.l)
			case "j":
				d.row = increment(d.row, d.height)
			case "k":
				d.row = decrement(d.row, d.m.t)
			case "l":
				d.col = increment(d.col, d.width)
			case "i":
				b.Insert('i')
			}
			b.SetIndex(d)
		}
	}
}
