package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:4321")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}

		go func(c net.Conn) {
			defer func() {
				c.Close()
			}()

			payload, err := Decode (c) 

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("Client : %s\n", string(payload.Byte()))

			p := Binary("Receive message : " + string(payload.Byte()))

			_, err = p.WriteTo(c)

			if err != nil {
				fmt.Println(err)
				return
			}
		}(conn)
	}
}
