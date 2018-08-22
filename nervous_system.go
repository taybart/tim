package main

/*
BusMessage should be a part of a class for subscribing
loadbalancer: should emulate socket.io, also have a socket interface
							 subscribe(), send(), emit()
*/
import (
	"github.com/taybartski/log"
)

// bus message
type BusMessage struct {
	nonce   int
	message string
	caller  string
}

func (parent *BusMessage) send(bus chan BusMessage, message string) {
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
	channels   []string
	callers    []string
	SpinalCord chan interface{}
}

func (ns *NervousSystem) NewChannel(chanName string) {
	ns.channels = append(ns.channels, chanName)
}

func (ns *NervousSystem) IsValidCaller(bm BusMessage) bool {
	for _, c := range ns.callers {
		if c == bm.caller {
			return true
		}
	}
	return false
}

func WakeUp() NervousSystem {
	return NervousSystem{}
}

func (ns *NervousSystem) AutoNom() {
	for {
	}
}

type Finger struct {
	nerve *net.Listener
}

func (ns *NervousSystem) AddFinger() {
	l, err := net.Listen("tcp", ":13130")
	if err != nil {
		log.Error("listen error:", err)
	}
	defer l.Close()

	subs := make([]chan string, 0)
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

	for {
		select {
		case data := <-input:
			// bus <- data
			blast(data, subs)
		case fd := <-newConns:
			if fd == nil {
				panic("Connection error")
			}
			bus := make(chan string)
			subs = append(subs, bus)
			go relay(fd, bus)
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

// Socket Server
func (f Finger) feel(c net.Conn, bus chan string) {
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
func (f Finger) touch(c net.Conn, bus chan string) {
	for {
		data := <-bus
		_, err := c.Write([]byte(data))
		if err != nil {
			log.Error("Write: ", err)
		}
	}
}

func (f Finger) blast(data string, chans []chan string) {
	for _, ch := range chans {
		ch <- data
	}
}
