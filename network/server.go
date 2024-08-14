package network

import (
	"encoding/binary"
	"io"
	"net"
	"os"
	"strconv"
	"log"
	"lesgofile/settings"
)

//wait for client and download file

func Reciver(port string) {
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
			go saveFile(conn)
		}
	}
}

func saveFile(conn net.Conn) {
	defer conn.Close()

	var size int64
	var n_read int
	dim, err := strconv.Atoi(settings.SETTINGS["DIM_BUFFER"])
	if err != nil {
		log.Panic("could not load buffer dimension" )
	}
	buffer := make([]byte, dim)
	filename := make([]byte, dim)

	binary.Read(conn, binary.BigEndian, &size)

	//read filename
	n_read, err = conn.Read(filename[:size])
	if err != nil {
		log.Panic("could not read filename" )
	}
	log.Print( "recived "+string(filename[:n_read])+" filename\n")

	//create file
	file, err := os.Create(settings.SETTINGS["DESTINATION_FOLDER"] + string(filename[:n_read]))
	if err != nil {
		log.Panicf("could not open %s for writing",settings.SETTINGS["DESTINATION_FOLDER"] + string(filename[:n_read]))
	}
	defer file.Close()

	//write file
	for err != io.EOF {
		n_read, err = conn.Read(buffer)
		if err != nil {
			log.Panicf("error in reading from %s",conn.RemoteAddr())
		}
		file.Write(buffer[:n_read])
	}

	log.Print("connection terminated")

}
