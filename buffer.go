package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/taybartski/log"
	"io/ioutil"
)

// HandleBuffer Used for navigating buffers
func HandleBuffer(fileName string, bus chan string) {
	file, err := ioutil.ReadFile(fileName)
	check(err)
	log.Info(string(file))

	// fd, err := os.Open(fileName)
	// check(err)

	index := 0
	print(0, 3, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("%s", string(file[index:index+500])))
	termbox.Flush()
	for {
		select {
		case s := <-bus:
			switch s {
			case "h":
				// log.Debug("left: index %v", index)
				index--
				if index < 0 {
					index = 0
				}
			case "l":
				// log.Debug("right: index %v", index)
				index++
				if index > len(file) {
					index = len(file) - 1
				}
			}
			log.Infoln(string(file[index : index+3]))
			print(0, 3, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("%s", string(file[index:index+500])))
		}
		termbox.Flush()
		// log.Debug("index %v", index)

		// buf := make([]byte, 5)
		// _, err := fd.Read(buf)
		// check(err)
		// log.Info("buf: %s", s)
		// fmt.Printf("%s\r", string(buf))
	}
}

// ombud testid: OMB-0411-SCLQ
