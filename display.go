package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	// "math"
)

type Display struct {
	row    int
	col    int
	width  int
	height int
	buf    []byte
}

func (d Display) Draw() {
	d.DrawBuffer()
	d.DrawMode("normal")
	d.DrawCursor()
}

// DrawMode displays current modality
func (d Display) DrawMode(mode string) {
	lines := len(d.buf)/d.width + 1
	bar := fmt.Sprintf("%s < lines(%d), len(%d), width(%d)", mode, lines, len(d.buf), d.width)
	draw(d.width-len(bar), d.height-1, termbox.ColorRed, termbox.ColorDefault, bar)
}

// DrawBuffer displays current file
func (d Display) DrawBuffer() {
	editWidth := d.width - 3
	lines := (len(d.buf) / editWidth) + 1

	for i := 0; i < lines; i++ {
		draw(0, i, termbox.ColorYellow, termbox.ColorDefault, fmt.Sprintf("%d", i))
		draw(3, i, termbox.ColorDefault, termbox.ColorDefault, string(d.buf[i*editWidth:(i+1)*editWidth]))
	}
}
func (d Display) DrawCursor() {
	termbox.SetCursor(d.col, d.row)
}
