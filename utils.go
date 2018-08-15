package main

import (
	// "fmt"
	"github.com/nsf/termbox-go"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func draw(x, y int, fg, bg termbox.Attribute, s string) {
	for _, c := range s {
		termbox.SetCell(x, y, rune(c), fg, bg)
		x++
	}
}
func printCursor(x, y int) {
	termbox.SetCursor(x, y)
}
