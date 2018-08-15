package main

import (
// "github.com/nsf/termbox-go"
// "io/ioutil"
)

type Buffer struct {
	data  []byte
	index int
	mode  string
}

func (b *Buffer) Insert(letter byte) {
	b.data = append(b.data, 0)
	copy(b.data[b.index+1:], b.data[b.index:])
	b.data[b.index] = letter
}
func (b *Buffer) SetIndex(d Display) {
	b.index = (d.col - d.m.l) + (d.row-d.m.t)*(d.width-d.m.l-d.m.r)
	if b.index >= len(b.data) {
		b.index = len(b.data) - 1
	}
}
