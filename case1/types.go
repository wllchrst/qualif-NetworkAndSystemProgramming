package main

import (
	"encoding/binary"
	"fmt"
	"io"
)

const (
	BinaryType uint8 = iota + 1
	StringType
)

type Payload interface {
	io.WriterTo
	io.ReaderFrom
	Byte() []byte
	String() string
}

type Binary []byte

func (message Binary) Byte() []byte {
	return message
}

func (message Binary) String() string {
	return string(message)
}

func (message Binary) WriteTo(write io.Writer) (int64, error) {
	err :=  binary.Write(write, binary.BigEndian, BinaryType)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	var n int64 = 1

	err = binary.Write(write, binary.BigEndian, uint32(len(message)))

	if err != nil {
		fmt.Println(err)
		return 0, err
	}	

	n += 4

	output, err := write.Write(message)

	return n + int64(output), err
}

func (message *Binary) ReadFrom(read io.Reader) (int64, error) {
	var typ uint8

	err := binary.Read(read, binary.BigEndian, &typ)

	if err != nil { 
		fmt.Println(err)
		return 0, err
	}

	var n int64 = 1

	var size int32

	err = binary.Read(read, binary.BigEndian, &size)

	if err != nil { 
		fmt.Println(err)
		return 0, err
	}

	n += 4

	*message = make([]byte, size)

	o, err := io.ReadFull(read, *message)

	return n + int64(o), err
}

func Decode(read io.Reader) (Payload, error){ 
	var typ uint8

	err := binary.Read(read, binary.BigEndian, &typ)

	if err != nil { 
		fmt.Println(err)
		return nil, err
	}

	var size int32 

	err = binary.Read(read, binary.BigEndian, &size)

	if err != nil { 
		fmt.Println(err)
		return nil, err
	}

	payload := make(Binary, size)

	_, err = io.ReadFull(read, payload)

	if err != nil { 
		fmt.Println(err)
		return nil, err
	}

	return &payload, nil
}
