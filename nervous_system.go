package main

/*
BusMessage should be a part of a class for subscribing
loadbalancer: should emulate socket.io, also have a socket interface
							 subscribe(), send(), emit()
*/
import (
	"github.com/nsf/termbox-go"
	"github.com/taybartski/log"
	"net"
	"time"
)

// bus message
type NervePulse struct {
	nonce   int
	message string
	caller  string
}

func (parent *NervePulse) send(bus chan BusMessage, message string) {
	parent.nonce++
	bm := BusMessage{parent.nonce, message, parent.caller}
	if bm.caller == "" {
		log.Error("Caller not defined")
		// return "Caller not defined"
	}
	bm.message = message
	bus <- bm
}

type NervousSystem struct {
	//? channels []chan interface{}
	// channels   []string
	// callers    []string
	SpinalCord chan interface{} // Main bus
	Nerves     int              // Internal listeners
	Fingers    int              // External listeners
}

// func (ns *NervousSystem) NewChannel(chanName string) {
// ns.channels = append(ns.channels, chanName)
// }

// func (ns *NervousSystem) IsValidCaller(bm BusMessage) bool {
// for _, c := range ns.callers {
// if c == bm.caller {
// return true
// }
// }
// return false
// }

func WakeUp() NervousSystem {
	nerve := make(chan interface{})
	return NervousSystem{SpinalCord: nerve, Nerves: 0, Fingers: 0}
}

func (ns *NervousSystem) AutoNom() {
	// socketOutput := make(chan interface{})
	go ns.AddFingerListener()
autonom:
	for {
		select {
		case pulse := <-ns.SpinalCord:
			// log.Info("recieved: %v", pulse)
			ns.blast(pulse)
			if pulse.(string) == "q" {
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				break autonom
			}
		}
	}
}

func (ns *NervousSystem) AddNerve() {
	ns.Nerves++
}
func (ns *NervousSystem) AddFinger() {
	ns.Fingers++
}

func (ns *NervousSystem) AddFingerListener() {
	// ns.AddFinger()
	ns.AddNerve()
	l, err := net.Listen("tcp", ":13130")
	if err != nil {
		log.Error("listen error:", err)
	}
	defer l.Close()

	newConns := make(chan net.Conn)

	go func(l net.Listener) {
		for {
			c, err := l.Accept()
			if err != nil {
				newConns <- nil
				return
			}
			newConns <- c
		}
	}(l)

	subs := make([]chan string, 0)
	for {
		select {
		case data := <-ns.SpinalCord:
			// bus <- data
			blast(data.(string), subs)
		case fd := <-newConns:
			if fd == nil {
				panic("Connection error")
			}
			bus := make(chan string)
			subs = append(subs, bus)
			go write(fd, bus)
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

// type Finger struct {
// // nerve chan interface{}
// digit *net.Listener
// }

// Socket Server
func read(c net.Conn, bus chan string) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		println("Server got:", string(data))
	}
}
func write(c net.Conn, bus chan string) {
	for {
		data := <-bus
		_, err := c.Write([]byte(data))
		if err != nil {
			log.Error("Write: ", err)
		}
	}
}
func blast(data string, chans []chan string) {
	for _, ch := range chans {
		ch <- data
	}
}
func (ns NervousSystem) blast(data interface{}) {
	for i := 1; i < ns.Nerves; i++ {
		ns.SpinalCord <- data
	}
}
