package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	dial, err := net.Dial("tcp", "localhost:9999")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer dial.Close()

	payload := Binary("Hello, World\n")

	_, err = payload.WriteTo(dial)

	if err != nil {
		fmt.Println(err)
		return
	}

	dial.SetReadDeadline(time.Now().Add(5 * time.Second))

	p, err := Decode(dial)

	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			fmt.Println("Connection timeout")
		} else {
			fmt.Println(err)
		}
	}

	fmt.Println(p)
}
