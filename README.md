# TIM
Taylor's Vim!

This is in early stages, please feel free to make a PR or comment!

`mkdir -p $GOPATH/src/github.com/taybartski/ && cd $GOPATH/src/github.com/taybartski && git clone https://github.com/taybartski/tim && cd tim && go build`

`./tim test.txt`

## Concept

Tim is good at word processing, he has a nervous system that lets all parts of him communicate.

Users of Tim can add fingers to his body. This is done by adding a TCP listener and requesting/subscribing to info!

Here is an example of a finger in go:
```go
package main

import (
	"io"
	"net"
	"time"
)

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		println(string(buf[0:n]))
	}
}

func write(c io.Writer) {
	for {
		_, err := c.Write([]byte("hi"))
		if err != nil {
			break
		}
		time.Sleep(1e9)
	}
}

func main() {
	c, err := net.Dial("tcp", ":13130")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	reader(c)
}

```

## Structure
![image of tim](https://github.com/taybartski/tim/raw/master/static/tim.png)

## Topics/Paths/Channels
These are the current paths:

* "input"
* "display"
* "color"
* "buffer"

