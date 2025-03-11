package network

import (
	"encoding/binary"
	"io"
	"net"
	"os"
	"log"
)

func Reciver(port string,buffer_dimension int,destination_folder string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Panicf("unable to listen over port: %s", port)
	}
	defer ln.Close()
	log.Printf( "start listen on port: %s", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print("unable to accept request from client")
		} else {
			log.Printf( "connected to client %s", conn.RemoteAddr())
			go saveFile(conn,buffer_dimension,destination_folder)
		}
	}
}

func saveFile(conn net.Conn,buffer_dimension int,destination_folder string) {
	defer conn.Close()

	var size int64
	var n_read int
	buffer := make([]byte, buffer_dimension)
	filename := make([]byte, buffer_dimension)

	binary.Read(conn, binary.BigEndian, &size)

	// read filename
  n_read, err := conn.Read(filename[:size])
	if err != nil {
		log.Print("could not read filename" )
	}
	log.Printf("recived "+string(filename[:n_read])+" filename\n")

	// create file
	file, err := os.Create(destination_folder + string(filename[:n_read]))
	if err != nil {
		log.Printf("could not open %s for writing",destination_folder + string(filename[:n_read]))
	}
	defer file.Close()

	// write file
	for err != io.EOF {
		n_read, err = conn.Read(buffer)
		if err != nil && err != io.EOF{
			log.Printf("error in reading from %s",conn.RemoteAddr())
			log.Printf(err.Error())
		}
		file.Write(buffer[:n_read])
	}

	log.Printf("connection terminated")

}
