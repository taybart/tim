package main

import (
// "github.com/nsf/termbox-go"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func increment(counter, max int) int {
	counter++
	if counter > max {
		counter = max - 1
	}
	return counter
}

func incrementBy(counter, by, max int) int {
	counter += by
	if counter > max {
		counter = max - 1
	}
	return counter
}

func decrement(counter, min int) int {
	counter--
	if counter < min {
		counter = min
	}
	return counter
}

func decrementBy(counter, by, min int) int {
	counter -= by
	if counter < min {
		counter = min
	}
	return counter
}
